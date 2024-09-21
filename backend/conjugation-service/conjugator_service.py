from mlconjug3 import Conjugator

class ConjugatorService:
    def get_conjugation(self, verb, tense="presente", person="1s"):
        # Initialize conjugator for Italian verbs
        conjugator = Conjugator(language='it')

        # Get the conjugation for the verb
        verb_conjugation = conjugator.conjugate(verb, subject='abbrev')
        print("WHAT")
        print(verb_conjugation.iterate())
        # Check if the tense and person are available
        try:
            # For Italian: 'indicativo' and 'presente' are typical structures for verb conjugation.
            conjugation_form = verb_conjugation['Indicativo']['Indicativo presente']
            return conjugation_form
        except KeyError:
            # Return an error message if the conjugation form is not found
            return f"Invalid conjugation form for {verb} in tense {tense} and person {person}"
