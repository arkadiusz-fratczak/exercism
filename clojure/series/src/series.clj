(ns series)

(defn slices [string length]
  (cond
    (>= 0 length) [""]
    :else (for [s (range (- (count string) length -1))]
            (subs string s (+ s length))))
  )
