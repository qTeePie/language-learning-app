# Dictionary API language learning app

Translation happens visa Lexica API
Source Dictionary Password (english) &rarr; target language. </br>
Output is checked against Lexica Multigloss Dictionary to limit senses to most frequent with correct POS (place of speech).

Verb Conjugation Possible APIs / packages
https://api.verbix.com/conjugator/html

Python:
https://pypi.org/project/mlconjug3/

Docker image apertium for verb conjagulations & verb POS
Github Repo: https://github.com/apertium/apertium-apy

## Services

### Management Service

Per now for admins of page. Each timeperiod is assigned words (noun / verb or other POS) for the user to learn.
This service manages the flow of learning on the app (as presented to user).

#### Endpoints

### Translation Service

API calls to free translation APIs to fetch translations for given word.

#### Endpoints
