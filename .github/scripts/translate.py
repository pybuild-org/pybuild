import os
import pathlib
from openai import OpenAI
from concurrent.futures import ThreadPoolExecutor

AI_API_KEY = os.getenv("AI_API_KEY")
AI_API_URL = os.getenv("AI_API_URL")
AI_MODEL = os.getenv("AI_MODEL")

client = OpenAI(
    api_key=AI_API_KEY,
    base_url=AI_API_URL,
)

docs_dir = pathlib.Path("docs")
doc_base_name = "base.md"
doc_base = docs_dir / doc_base_name
doc_base_file = doc_base.read_text(encoding="UTF-8")

for item in docs_dir.iterdir():
    if item.is_file() and item.name != doc_base_name:
        print("remove old", item.name)
        item.unlink()

TRANSLATE_TARGET = [
    "zh-CN",
    "zh-TW",
    "en-US",
    "ja-JP",
    "ru-RU",
]


def translate_text(base: str, target: str) -> str:
    developer_instruction = " ".join(
        [
            "You are a professional translator.",
            f"Translate the incoming text directly into {target}.",
            f"Maintain the original tone, formatting, and paragraphs.",
            "Do not include any introductory or concluding remarks—return only the translated text.",
        ]
    )

    response = client.chat.completions.create(
        model=AI_MODEL,  # type: ignore
        temperature=0.3,
        messages=[
            {"role": "developer", "content": developer_instruction},
            {"role": "user", "content": f"Text to translate:\n{base}"},
        ],
    )

    content = response.choices[0].message.content
    if not content:
        return ""

    return content.strip()


def process_translation(target: str):
    print("translate to", target)
    content = translate_text(doc_base_file, target)

    target_file = docs_dir / (target + ".md")
    target_file.write_text(content, encoding="UTF-8")


with ThreadPoolExecutor(max_workers=3) as executor:
    executor.map(process_translation, TRANSLATE_TARGET)
