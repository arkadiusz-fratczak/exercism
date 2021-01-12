(ns bob
  (:require [clojure.string :as str]))

(declare contains-letter? is-upper-cased? ends-with-qst? does)

(defn response-for [s]
  (let [sentence (str/trim s)]
    (cond
      (does sentence [contains-letter? is-upper-cased? ends-with-qst?])
      "Calm down, I know what I'm doing!"
      (does sentence [contains-letter? is-upper-cased?])
      "Whoa, chill out!"
      (ends-with-qst? sentence)
      "Sure."
      (str/blank? sentence)
      "Fine. Be that way!"
      :else
      "Whatever.")))

(defn contains-letter? [s]
  (some? (some #(Character/isLetter %) s)))

(defn is-upper-cased? [s]
  (= (str/upper-case s) s))

(defn ends-with-qst? [s]
  (= \? (last s)))

(defn does [s, requirements]
  (every? #(% s) requirements))
