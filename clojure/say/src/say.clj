(ns say
  (:require [clojure.string :as str]))

(def spelling {0 [""]
               1 ["one", "eleven", "ten"]
               2 ["two", "twelve", "twenty"]
               3 ["three", "thirteen", "thirty"]
               4 ["four", "fourteen", "forty"]
               5 ["five", "fifteen", "fifty"]
               6 ["six", "sixteen", "sixty"]
               7 ["seven", "seventeen", "seventy"]
               8 ["eight", "eighteen", "eighty"]
               9 ["nine", "nineteen", "ninety"]})

(def scales ["" "thousand" "million" "billion" "trillion"])

(defn units [u] (get-in spelling [u 0]))

(defn teens [t] (get-in spelling [t 1]))

(defn tens [t] (get-in spelling [t 2]))

(defn spell-chunk [num]
  (let [h (quot num 100)
        hrem (rem num 100)
        t (quot hrem 10)
        u (rem num 10)]
    (cond
      (and (> h 0) (= t 0) (= u 0)) (str (units h) " hundred")
      (not= h 0) (str (units h) " hundred " (spell-chunk hrem))
      (= t 0) (units u)
      (= u 0) (tens t)
      (= t 1) (teens u)
      :else (str (tens t) "-" (units u)))))

(defn partition-by-thousands [num]
  (reverse (map (fn [n scale] [(mod n 1000) scale])
                (take-while pos? (iterate #(quot % 1000) num))
                scales)))

(defn chunks [num]
  (loop [nr num
         acc []]
    (let [q (quot nr 1000)
          r (rem nr 1000)]
      (if (= q 0)
        (cons r acc)
        (recur q (cons r acc))))))

(defn chunk-with-scale [chunk, scale]
  (if-not (= chunk 0)
    [(spell-chunk chunk) scale]))

(defn spell-number [num]
  (let [ch (chunks num)
        q (count ch)
        s (reverse (take q scales))]
    (str/join " " (filter #(not (str/blank? %)) (flatten (map chunk-with-scale ch s))))))

(defn number [num]
  (cond
    (= num 0) "zero"
    (< num 0) (throw (IllegalArgumentException. "number must be positive"))
    (> num 999999999999) (throw (IllegalArgumentException. "number must be below 1000000000000"))
    :else (spell-number num)))

