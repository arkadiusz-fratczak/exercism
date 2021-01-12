(ns anagram
  (:require [clojure.string :as str]))

(defn is-anagram? [word anagram]
  (let [lower-cased-word (str/lower-case word)
        lower-cased-anagram (str/lower-case anagram)]
    (if-not (= lower-cased-word lower-cased-anagram)
      (= (frequencies lower-cased-word) (frequencies lower-cased-anagram)))))

(defn anagrams-for [word prospect-list]
  (filter (partial is-anagram? word) prospect-list))


