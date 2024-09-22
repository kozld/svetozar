import sys, torch

# Load model directly
from transformers import AutoTokenizer, AutoModelForSequenceClassification

tokenizer = AutoTokenizer.from_pretrained("sismetanin/rubert-toxic-pikabu-2ch")
model = AutoModelForSequenceClassification.from_pretrained("sismetanin/rubert-toxic-pikabu-2ch")

# Текст для классификации
text = sys.argv[1]

# Токенизация текста
inputs = tokenizer(text, return_tensors="pt", padding=True, truncation=True)

# Прогнозирование на основе модели
with torch.no_grad():
    outputs = model(**inputs)
    logits = outputs.logits

# Получение предсказанного класса
predicted_class = torch.argmax(logits, dim=1).item()

print(f"Предсказанный класс: {predicted_class}")
