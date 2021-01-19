(ns clock)

(def minutes-in-day (* 24 60))

(defn- hours [clock]
  (mod (quot (:m clock) 60) 24))

(defn- minutes [clock]
  (mod (:m clock) 60))

(defn- normalize-minutes [minutes]
  (mod minutes minutes-in-day))

(defn clock->string [clock]
  (format "%02d:%02d" (hours clock) (minutes clock))
)

(defn clock [hours minutes]
  {:m (normalize-minutes (+ (* 60 hours) minutes))}
)

(defn add-time [clock time]
  {:m (normalize-minutes (+ time (:m clock)))}
)
