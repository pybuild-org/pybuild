import os
import pathlib
import openai
import concurrent.futures

AI_API_KEY = os.getenv("AI_API_KEY")
AI_API_URL = os.getenv("AI_API_URL")
AI_MODEL = os.getenv("AI_MODEL")

client = openai.OpenAI(
    api_key=AI_API_KEY,
    base_url=AI_API_URL,
)

docs_dir = pathlib.Path("docs")
docs_base_dir = docs_dir / "base"

TRANSLATE_TARGET = [
    "zh-CN",
    "zh-TW",
    "en-US",
    "ja-JP",
    "ru-RU",
]


def main():
    clean()
    translate()


def clean():
    for file_path in docs_dir.rglob("*"):
        if file_path.is_file():
            if docs_base_dir not in file_path.parents and file_path != docs_base_dir:
                file_path.unlink()
                print("remove", file_path.absolute())

    for dir_path in sorted(docs_dir.rglob("*"), reverse=True):
        if dir_path.is_dir():
            if docs_base_dir not in dir_path.parents and dir_path != docs_base_dir:
                try:
                    dir_path.rmdir()
                    print("remove", dir_path.absolute())
                except:
                    pass


def translate():
    tasks: list[tuple[str, str, pathlib.Path]] = []
    for file_path in docs_base_dir.rglob("*.md"):
        if file_path.is_file():
            source_text = file_path.read_text(encoding="utf-8")
            relative_path = file_path.relative_to(docs_base_dir)

            for target in TRANSLATE_TARGET:
                tasks.append((source_text, target, relative_path))

    with concurrent.futures.ThreadPoolExecutor(max_workers=3) as executor:
        for source, target, rel_path in tasks:
            executor.submit(translate_single_target, source, target, rel_path)


def translate_single_target(source_text: str, target: str, relative_path: pathlib.Path):
    target_file_path = docs_dir / target / relative_path
    target_file_path.parent.mkdir(parents=True, exist_ok=True)

    print("translate", target_file_path.absolute())
    translated_text = translate_text(source_text, target)
    target_file_path.write_text(translated_text, encoding="utf-8")


def translate_text(source: str, target: str) -> str:
    system_prompt = "\n".join(
        [
            "You are a professional, expert translator.",
            f"Your task is to translate the user's input text into the target language: '{target}'.",
            "Strictly follow these rules:",
            "1. Maintain the exact original Markdown formatting (headers, lists, code blocks, links, bold, italic, etc.).",
            "2. Do not translate code snippets, variable names, URLs, or technical terms that should remain intact.",
            "3. Deliver ONLY the translated content. Do not include any explanations, greetings, or extra commentary.",
        ]
    )

    response = client.chat.completions.create(
        model=AI_MODEL,  # type: ignore
        temperature=0.3,
        messages=[
            {"role": "system", "content": system_prompt},
            {"role": "user", "content": source},
        ],
    )

    return response.choices[0].message.content  # type: ignore


if __name__ == "__main__":
    main()
