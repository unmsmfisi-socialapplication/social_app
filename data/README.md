# App Social - Data
# Python para desarrollo de las distintas caracteristicas que requiera el sistema , formato
jupyter para el testeo inicial

## Estructura de Carpeta

  ```
├── cloud_platorm
│   ├── utilities
│   └── client_secrets.json
├── models_testing
│   ├── post_classification
│   ├── sentiment_analysis
│   └── spam_detector
│
├── post_classification
│   └──post_classification.py
├── sentiment_analysis
│   └──sentiment_analysis.py
├── spam_detector
│   └──spam_detector.py
 ```

### Cloud_platorm

En la carpeta de Cloud_platform iran todos los componentes requeridos para configurar el almacenamiento
en la nube

#Models_testing
En esta carpeta va todo el testeo de modelos para seleccionar aquellos que obtengan mejores metricas en las pruebas
del equipo , cada una de las subcarpetas representa una caracteristica que se asigno al equipo , esta hecho para almacenar
archivos de jupyter notebook en base a los caracteristicas

# postpost_classification / sentiment_analysis /spam_detector
Son carpetas que deben de almacenar los script de python oficiales de los modelos , estos netamente
se agregar al finalizar las pruebas en el ambiente de Jupyter notebook

#Entorno local
Se definen en este todas las variables que procederan a ser usadas en el desarrollo , evitando su exposicion
publica al repositorio .Considerar que python usa la libreria os para el manejo del entorno: 

#Lee la ubicación del archivo de credenciales desde la variable de entorno
credentials_path = os.environ.get('tuta_entorno')
