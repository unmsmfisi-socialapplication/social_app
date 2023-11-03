import requests
import csv

def consume_api_and_save_csv(api_url):

    # Make a GET request to the API for information
    response = requests.get(api_url)

    # Check if the application was successful
    if response.status_code == 200:
        data = response.json()

        if data:
            # Get the field names of the first object in the JSON response
            fieldnames = data[0].keys()

            # Use the hostname as the name of the CSV file
            csv_filename = api_url.split("//")[-1].replace("/", "_") + ".csv"

            # Create a CSV file and enter the data
            with open(csv_filename, "w", newline="") as csvfile:
                writer = csv.DictWriter(csvfile, fieldnames=fieldnames)

                # Write the header 
                writer.writeheader()

                # Iterate through the data and write it to the CSV file
                for item in data:
                    writer.writerow(item)

            print(f"Los datos se han guardado en el archivo '{csv_filename}'.")
        else:
            print("No se encontraron datos en la respuesta del API.")
    else:
        print("Error al obtener datos del API. Código de estado:", response.status_code)

    return True
 

# Calling the function with the API URL 
api_url = "https://jsonplaceholder.typicode.com/posts" # Replace with the URL of the API you want to consume
consume_api_and_save_csv(api_url)
