from flask import Flask, render_template, request, redirect, url_for
import json

import joblib
import boto3

app = Flask(__name__)

@app.route('/', methods=['GET', 'POST'])
def index():
    #return "Hello World"
    if request.method == 'POST':
        dia = request.form['dia']
        hora = request.form['hora']
        mes = request.form['mes']

        data = {'Dia': dia, 'Hora': hora, 'Mes': mes}
        #data = [float(area), float(perimetro), float(cantidad_vertices), float(nivel), float(rmc_largo), float(rmc_alto), float(rmc_ratio)]

        return redirect(url_for('prediccion', data=json.dumps(data)))

    return render_template('input_form.html')

@app.route('/prediccion<data>')
def prediccion(data):
    # Ruta modelo para la prediccion en Local
    model = joblib.load('model.joblib')

    formated_data = json.loads(data)

    bien_data = [int(formated_data.get('Dia')), int(formated_data.get('Hora')), int(formated_data.get('Mes'))]

    bien_data_for_endpoint = ",".join(str(x) for x in bien_data)
    bien_data_for_endpoint = bien_data_for_endpoint + "\n" + bien_data_for_endpoint

    print(bien_data)
    print(bien_data_for_endpoint)

    # Predicción con Endpoint SageMaker
    endpoint = 'sagemaker-scikit-learn-2022-12-06-01-34-50-163'
    client = boto3.client('sagemaker-runtime', 'us-east-1')

    response = client.invoke_endpoint(EndpointName=endpoint, Body=bien_data_for_endpoint, ContentType='text/csv')

    #print(response)

    # Devuelve un string con el valor de la predicción, entonces extraemos el segundo caracter que contiene el valor de la predicción y lo convertimos a entero
    prediccion = int(response['Body'].read().decode()[1:2])


    print(prediccion)

    # Predicción con modelo local
    #prediccion = model.predict([bien_data])


    complemento = f"El resultado del servcicio: {bien_data} es"
    respuesta = f"{complemento} NOTIFICAR" if prediccion[0] == 0 else f"{complemento} NO NOTIFICAR"

    return respuesta


if __name__ == "__main__":
    app.run(debug=True, host='0.0.0.0', port=5000)