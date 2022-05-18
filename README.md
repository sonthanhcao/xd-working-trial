# Note
# Buddy Work

## Optimize Current
Costing
Slowing Pipeline build, Slowing CI Server

Some Common reason:
- Reach Concurrency limit of Buddy Work Runner/Worker
- Slow Pipeline Build:
    + Missing Cache
    + Not ennough Resource Usage
    + Pending Build due to concurrency capability of Buddy Worker


Possible Action:
- Develop solution for deploying new buddy worker at fastest way vm base image for deploying buddy worker 
- Develop custom watcher and monitor concurrency build of Buddy.Work then trigger autoscale buddy worker VM so we can scale on demand
- Develop custom schedule job to scale based on high peak time
- 


## Futrure
- Migrate to Gitlab and Gitlab CI, or at least CI stuff
- Source Code and 