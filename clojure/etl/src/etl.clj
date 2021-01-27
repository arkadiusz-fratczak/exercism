(ns etl
  (:require [clojure.string :as str]))

(defn transform [source]
  (reduce-kv (fn [m k v]
               (merge m (reduce (fn [im ik] (assoc im (str/lower-case ik) k)) {} v)))
             {} source))
