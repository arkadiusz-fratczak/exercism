(ns nucleotide-count)

(defn nucleotide-counts [strand]
  (reduce #(update %1 (first %2) + (count %2))
          {\A 0, \T 0, \C 0, \G 0}
          (partition-by identity strand)))

(defn count-of-nucleotide-in-strand [nucleotide strand]
  (when-not (contains? {\A 0, \T 0, \C 0, \G 0} nucleotide)
    (throw (IllegalArgumentException. "invalid nucleotide")))
  (get (nucleotide-counts strand) nucleotide 0))
