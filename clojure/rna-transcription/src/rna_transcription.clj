(ns rna-transcription)

(def transcription-map {\G "C" \C "G" \T "A" \A "U"})

(defn strict-get [map key]
  {:pre [(contains? map key)]}
  (get map key))

(defn to-rna [dna]
  (apply str (map #(strict-get transcription-map %) dna)))

