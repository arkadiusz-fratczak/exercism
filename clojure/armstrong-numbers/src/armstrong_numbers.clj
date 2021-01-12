(ns armstrong-numbers)
(declare digits exp)

(defn armstrong? [num]
  (let [d (digits num)
        q (count d)]
    (= num (reduce (fn [a b] (+ a (exp b q))) 0 d))))

(defn exp [x n]
  (reduce * (repeat n x)))

(defn digits
  ([num] (digits num []))
  ([num acc]
   (let [quotient (quot num 10)
         reminder (mod num 10)]
     (if (= quotient 0)
       (conj acc reminder)
       (conj (digits quotient acc) reminder)))))
