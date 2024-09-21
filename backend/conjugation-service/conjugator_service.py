from mlconjug3 import Conjugator

# Define the mapping from pronouns to numeric values
pronoun_map = {
    "io": 1,
    "tu": 2,
    "egli/ella": 3,
    "noi": 4,
    "voi": 5,
    "essi/esse": 6
}

class ConjugatorService:
    def get_conjugation(self, verb, tense="presente", person="1s"):
        # Initialize conjugator for Italian verbs
        conjugator = Conjugator(language='it')

        # Get the conjugation
        verb_conjugation = conjugator.conjugate(verb, subject='pronoun')

        # Optional: Print conjugation structure
        # print(verb_conjugation.iterate())

        try:
            # Access the specific conjugation for 'Indicativo' and the given tense
            conjugation_form = verb_conjugation['Indicativo'][f'Indicativo {tense}']

            # Structure by pronoun order
            ordered_conjugation = {
                pronoun_map[pronoun]: f"{pronoun} {form}" for pronoun, form in conjugation_form.items()
            }
            
            return ordered_conjugation

        except KeyError:
            # Return an error message if the conjugation form is not found
            return f"Invalid conjugation form for {verb} in tense {tense} and person {person}"