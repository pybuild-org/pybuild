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


def translate_text(source: str, target: str) -> str:
    system = "\n".join(
        [
            "You are an expert, native-level professional translator.",
            f"Your task is to accurately translate the provided Markdown text into {target}.",
            "",
            "[CRITICAL RULES]",
            "1. Output ONLY the direct translation. Do NOT wrap the output in conversational text, and do NOT include any introductory or concluding remarks.",
            "2. Preserve all original Markdown formatting, syntax, line breaks, HTML tags, and code blocks exactly as they are. Do not translate code or configuration keys.",
            "3. Maintain the precise tone, style, and semantic meaning of the original text.",
            "4. The source text to translate is enclosed inside the <source_text> and </source_text> XML tags. Do NOT translate the tags themselves, and do NOT output these tags in your response.",
        ]
    )

    response = client.chat.completions.create(
        model=AI_MODEL,  # type: ignore
        messages=[
            {"role": "system", "content": system},
            {"role": "user", "content": f"<source_text>\n{source}\n</source_text>"},
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
