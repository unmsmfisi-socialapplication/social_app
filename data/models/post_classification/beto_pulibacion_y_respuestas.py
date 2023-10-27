# -*- coding: utf-8 -*-

#"!pip install transformers

# Importar
from textwrap import wrap
from transformers import AutoTokenizer, AutoModelForQuestionAnswering, pipeline
the_model = 'mrm8488/distill-bert-base-spanish-wwm-cased-finetuned-spa-squad2-es'
tokenizer = AutoTokenizer.from_pretrained(the_model, do_lower_case=False)
model = AutoModelForQuestionAnswering.from_pretrained(the_model)

print(model)

# Ejemplo tokenización
contexto = 'Yo soy Luis'
pregunta = '¿cómo me llamo?'

encode = tokenizer.encode_plus(pregunta, contexto, return_tensors='pt')
input_ids = encode['input_ids'].tolist()
tokens = tokenizer.convert_ids_to_tokens(input_ids[0])
for id, token in zip(input_ids[0], tokens):
  print('{:<12} {:>6}'.format(token, id))
  print('')

# Ejemplo de inferencia (pregunta-respuesta)
nlp = pipeline('question-answering', model=model, tokenizer=tokenizer)
salida = nlp({'question':pregunta, 'context':contexto})
print(salida)


def pregunta_respuesta(model, nlp):

    # Solicitar el contexto al usuario
    print('Por favor, ingrese el contexto:')
    print('-----------------')
    contexto = str(input())

    # Imprimir el contexto
    print('\nPublicacion:')
    print('-----------------')
    print('\n'.join(wrap(contexto)))

    # Definir la pregunta en el código
    pregunta = "¿De qué tema se habla?"

    # Imprimir la pregunta
    print('\nPregunta:')
    print('-----------------')
    print(pregunta)

    # Obtener respuesta
    salida = nlp({'question': pregunta, 'context': contexto})

    # Imprimir respuesta
    print('\nRespuesta:')
    print('-----------------')
    print(salida['answer'])

# Ejemplo de uso
# Llamada a la función sin proporcionar la pregunta
pregunta_respuesta(model, nlp)
