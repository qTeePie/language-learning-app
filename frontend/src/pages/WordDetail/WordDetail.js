import React from "react";
import { useParams } from "react-router-dom";
import { Header, Footer } from "components";
import styles from "./styles.module.css";

export const WordDetail = () => {
  const { word } = useParams();

  // Simulate fetching word details
  const wordDetails = {
    word: "Example",
    type: "Noun",
    definition: "A representative form or pattern.",
    plural: "Examples",
    conjugations: {
      present: [
        "I example",
        "You example",
        "He/She examples",
        "We example",
        "You (plural) example",
        "They example",
      ],
    },
    synonyms: ["sample", "instance", "illustration"],
    examples: [
      "This is an example of good practice.",
      "She set a good example for her younger siblings.",
    ],
  };

  return (
    <div className={styles.wordDetail}>
      <Header />
      <main className={styles.mainContent}>
        <h1>{wordDetails.word}</h1>
        <p>
          <strong>Type:</strong> {wordDetails.type}
        </p>
        <p>
          <strong>Definition:</strong> {wordDetails.definition}
        </p>
        {wordDetails.type === "Noun" && (
          <p>
            <strong>Plural:</strong> {wordDetails.plural}
          </p>
        )}
        {wordDetails.type === "Verb" && (
          <div>
            <h2>Conjugations</h2>
            <ul>
              {wordDetails.conjugations.present.map((form, index) => (
                <li key={index}>{form}</li>
              ))}
            </ul>
          </div>
        )}
        <h2>Synonyms</h2>
        <ul>
          {wordDetails.synonyms.map((synonym, index) => (
            <li key={index}>{synonym}</li>
          ))}
        </ul>
        <h2>Examples of Usage</h2>
        <ul>
          {wordDetails.examples.map((example, index) => (
            <li key={index}>{example}</li>
          ))}
        </ul>
      </main>
      <Footer />
    </div>
  );
};
