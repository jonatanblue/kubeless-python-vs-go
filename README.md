# kubeless-python-vs-go
Toy project to test deploying Go and Python functions with dependencies on Kubeless.

# Install

1. Install minikube
2. Install kubeless
3. Clone this repo

# Run

No functions created yet:
```
$ kubeless function ls
NAME	NAMESPACE	HANDLER	RUNTIME	DEPENDENCIES	STATUS
```

## Python

Deploy
```
$ kubeless function deploy python-loop --runtime python3.7 --handler loop.hello --from-file python/loop.py --dependencies python/requirements.txt
INFO[0000] Deploying function...
INFO[0000] Function python-loop submitted for deployment
INFO[0000] Check the deployment status executing 'kubeless function ls python-loop'
```

Show status. You may need to wait a few minutes before you see 1/1 READY, because the image needs to be pulled down.
```
$ kubeless function ls python-loop
NAME       	NAMESPACE	HANDLER   	RUNTIME  	DEPENDENCIES	STATUS
python-loop	default  	loop.hello	python3.7	flask       	1/1 READY
```

If it gets stuck for a long time, you can use kubectl to check the status of the pod. Here showing the interesting bits of the output.
```
$ kubectl describe pod $(kubectl get pods | grep python-loop | cut -d\  -f1)
...
Conditions:
  Type              Status
  Initialized       True
  Ready             True
  ContainersReady   True
  PodScheduled      True
...
Events:
  Type    Reason     Age   From               Message
  ----    ------     ----  ----               -------
  Normal  Scheduled  10m   default-scheduler  Successfully assigned default/python-loop-589d6fbc98-t2rlp to minikube
  Normal  Pulled     10m   kubelet, minikube  Container image "kubeless/unzip@sha256:f162c062973cca05459834de6ed14c039d45df8cdb76097f50b028a1621b3697" already present on machine
  Normal  Created    10m   kubelet, minikube  Created container
  Normal  Started    10m   kubelet, minikube  Started container
  Normal  Pulled     10m   kubelet, minikube  Container image "python:3.7" already present on machine
  Normal  Created    10m   kubelet, minikube  Created container
  Normal  Started    10m   kubelet, minikube  Started container
  Normal  Pulling    10m   kubelet, minikube  pulling image "kubeless/python@sha256:dbf616cb06a262482c00f5b53e1de17571924032e0ad000865ec6b5357ff35bf"
  Normal  Pulled     10m   kubelet, minikube  Successfully pulled image "kubeless/python@sha256:dbf616cb06a262482c00f5b53e1de17571924032e0ad000865ec6b5357ff35bf"
  Normal  Created    10m   kubelet, minikube  Created container
  Normal  Started    10m   kubelet, minikube  Started container
```
When Ready is True it's up and running.

Call
```
$ kubeless function call python-loop
Hello world!
```
Yay! It worked.

Check the logs for timing
```
$ kubeless function logs python-loop
...
2019-03-16 15:16:25.622213
1000000
2019-03-16 15:16:25.695977
...
```

## Golang

Deploy
```
$ kubeless function deploy go-loop --runtime go1.10 --handler loop.Hello --from-file golang/loop.go --dependencies golang/Gopkg.toml
INFO[0000] Deploying function...
INFO[0000] Function go-loop submitted for deployment
INFO[0000] Check the deployment status executing 'kubeless function ls go-loop'
```

Show status and wait until Status is 1/1 READY.
```
$ kubeless function ls go-loop
NAME   	NAMESPACE	HANDLER   	RUNTIME	DEPENDENCIES                         	STATUS
go-loop	default  	loop.Hello	go1.10 	                                     	1/1 READY
       	         	          	       	[[constraint]]
       	         	          	       	  name = "github.com/sirupsen/logrus"
       	         	          	       	  branch = "master"
```

Call
```
$ kubeless function call go-loop
Hello world!
```

Check the logs for timing
```
$ kubeless function logs go-loop
...
2019-03-16 15:29:41.516867695 +0000 UTC m=+111.692864650
1000000
2019-03-16 15:29:41.517866371 +0000 UTC m=+111.693863356
...
```

As you can see, Golang is way faster ;)
