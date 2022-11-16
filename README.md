# **README**

## 🚑🚑🚑🚑 Patient bracelet

## Project description
#### The demonstration program "**Patient bracelet**" estimates the risk of a heart attack by reading the necessary data from the sensors on a person's bracelet. Seven sensors together give a signal about the risk of a heart attack, warning with a message the person himself, the doctor whose person is under observation, and the person who is selected as an emergency contact.

## Features
1. [ ] fixing groundImpact
2. [ ] calculating body position (horizontally/vertically)
3. [ ] calculating number of steps
4. [ ] fixing heart rate
5. [ ] fixing body temperature
6. [ ] calculating hypertension
7. [ ] fixing glucose

## 📀 Preparing to launch a project

1. Clone project (need to push Code button and choose in what way you want to get project)
2. Open with IntelliJ IDEA (if you do not have this program - download it [here](https://www.jetbrains.com/idea/download/#section=mac))
3. You will also need Docker to this project. Download it from [here](https://www.docker.com)
4. You have to SET UP ROOT. How to do it you find [here](https://www.jetbrains.com/help/idea/configuring-goroot-and-gopath.html)


## 📌 How the project works

1. open terminal/cmd and put `docker run -d -p 6831:6831/udp -p 16686:16686 jaegertracing/all-in-one:latest`
2. Find  [service.go](exercise3/service/service.go) and press Run button or ^R in Intellij IDEA to start server
3. Open new tab in browser and type http://localhost:16686
4. Push button FIND TRACES. In Service menu find calc_risk_heart_attack and push FIND TRACES
5. If you need more information about JAEGER visit [site](https://www.jaegertracing.io) and read documentation# opentrace_go_example
# patient_bracelet
