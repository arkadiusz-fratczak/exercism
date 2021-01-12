(ns isbn-verifier)

(def isbn-regex #"\b(\d{9}[\dX]{1}|\d{1}-\d{3}-\d{5}-[\dX]{1})\b")

(defn to-number [ch]
  (cond
    (Character/isDigit ch) (Character/digit ch 10)
    (= \X ch) 10
    :else nil))

(defn parse [isbn]
  (filter some? (map to-number isbn)))

(defn isbn-sum [numbers]
  (apply + (map * (range 10 0 -1) numbers)))

(defn isbn? [isbn]
  (true?
    (if (re-matches isbn-regex isbn)
      (->> isbn
        parse
        isbn-sum
        (#(mod % 11))
        zero?))))
