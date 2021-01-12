(ns run-length-encoding)

(defn letter-reducer
  [acc-vector, letter]
  (if (and (last acc-vector) (clojure.string/ends-with? (last acc-vector) letter))
    (update acc-vector (dec (count acc-vector)) str letter)
    (conj acc-vector letter)))

(defn chunks [letters]
  (reduce letter-reducer [] letters))

(defn encode-chunk [chunk]
  (if (> 2 (count chunk))
    chunk
    (str (count chunk) (last chunk))))

(defn run-length-encode
  "encodes a string with run-length-encoding"
  [plain-text]
  (let [letters (clojure.string/split plain-text #"")]
    (apply str (map encode-chunk (chunks letters)))))

(defn is-digit [c]
  (and c (Character/isDigit c)))

(defn symbol-reducer
  [acc-vector, symbol]
  (if (is-digit (last (last acc-vector)))
    (update acc-vector (dec (count acc-vector)) str symbol)
    (conj acc-vector symbol)))

(defn cipher-chunks [symbols]
  (reduce symbol-reducer [] symbols))

(defn to-number [char-array]
  (if (empty? char-array)
    1
    (Integer/parseInt (apply str char-array))))

(defn decode-chunk [chunk]
  (let [symbols (clojure.string/split chunk #"")
        n (to-number (pop symbols))
        l (str (last symbols))]
    (apply str (repeat n l))))

(defn run-length-decode
  "decodes a run-length-encoded string"
  [cipher-text]
  (let [symbols (clojure.string/split cipher-text #"")]
    (apply str (map decode-chunk (cipher-chunks symbols)))))
