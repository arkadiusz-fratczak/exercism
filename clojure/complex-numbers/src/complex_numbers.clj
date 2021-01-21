(ns complex-numbers)

(defn real [[a b]] ;; <- arglist goes here
  a
)

(defn imaginary [[a b]] ;; <- arglist goes here
  b
)

(defn abs [[a b]] ;; <- arglist goes here
  (Math/sqrt (+ (* a a) (* b b)))
)

(defn conjugate [[a b]] ;; <- arglist goes here
  [a (- b)]
)

(defn add [[a b] [c d]] ;; <- arglist goes here
  [(+ a c) (+ b d)]
)

(defn sub [[a b] [c d]] ;; <- arglist goes here
  [(- a c) (- b d)]
)

(defn mul [[a b] [c d]] ;; <- arglist goes here
  [(- (* a c) (* b d)) (+ (* b c) (* a d))]
)

(defn div [[a b] [c d]] ;; <- arglist goes here
  [(double (/ (+ (* a c) (* b d)) (+ (* c c) (* d d))))
   (double (/ (- (* b c) (* a d)) (+ (* c c) (* d d))))]
)
