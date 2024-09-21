# main.py
from flask import Flask, jsonify, request
from conjugator_service import ConjugatorService

app = Flask(__name__)
conjugator_service = ConjugatorService()

# Define a route to fetch the verb conjugation
@app.route('/conjugate', methods=['GET'])
def conjugate():
    verb = request.args.get('verb', default="mangiare")  # Get the verb from query params
    tense = request.args.get('tense', default="presente")  # Optional tense parameter (default: presente)
    person = request.args.get('person', default="1s")     # Optional person parameter (default: 1s)
    
    # Get the specific conjugation form from the service
    conjugation_form = conjugator_service.get_conjugation(verb, tense, person)
    
    return jsonify({
        "verb": verb,
        "tense": tense,
        "person": person,
        "conjugation": conjugation_form
    })

if __name__ == '__main__':
    port = 5000
    print(f"Running on port {port}")
    app.run(debug=True, port=port)