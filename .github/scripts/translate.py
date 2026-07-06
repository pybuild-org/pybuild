import os

from openai import OpenAI

AI_API_KEY = os.getenv("AI_API_KEY")
AI_API_URL = os.getenv("AI_API_URL")
AI_MODEL = os.getenv("AI_MODEL")

client = OpenAI(
    api_key=AI_API_KEY,
    base_url=AI_API_URL,
)
