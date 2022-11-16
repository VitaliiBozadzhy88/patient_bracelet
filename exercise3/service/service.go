package main

import (
	"log"

	"exercises_go_tracing/exercise3/tracing"
	"github.com/opentracing/opentracing-go"
)

var Position bool
var Steps int
var GroundImpact bool
var HeartRate bool
var Temperature bool
var ArterialPressure bool
var Diabetes bool

var glucose = 6.1
var systolic = 180
var diastolic = 111

func main() {
	tracer, closer := tracing.Init("sensors_data_service")
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	CalcRiskAttack(Position, Steps, GroundImpact, HeartRate, Temperature, ArterialPressure, Diabetes)
}

func CalcRiskAttack(position bool, steps int, groundImpact bool, heartRate bool, temperature bool, arterialPressure bool, diabetes bool) {
	ParentSpan := opentracing.GlobalTracer().StartSpan("calc_risk_heart_attack")
	defer ParentSpan.Finish()

	position = GetHumanBodyPosition("horizontally", ParentSpan)
	steps = GetPedometer(0, ParentSpan)
	groundImpact = GetGroundImpact(true, ParentSpan)
	heartRate = GetHeartRate(101, ParentSpan)
	temperature = GetBodyTemperature(39, ParentSpan)
	arterialPressure = GetArterialPressure(systolic, diastolic, ParentSpan)
	diabetes = GetDiabetes(glucose, ParentSpan)

	ParentSpan.LogKV(
		"position", position,
		"ground impact", groundImpact,
		"heart rate", heartRate,
		"temperature", temperature,
		"arterial pressure", arterialPressure,
		"diabetes", diabetes,
	)

	if position == true &&
		steps == 0 &&
		groundImpact == true &&
		heartRate == true &&
		temperature == true &&
		arterialPressure == true &&
		diabetes == true {
		log.Println(MessageToPatient(ParentSpan), MessageToDoctor(ParentSpan), MessageToEmergencyContact(ParentSpan))
	} else {
		log.Println("You are OK!")
	}
}

// GetHumanBodyPosition calculates human body position - horizontally/vertically
func GetHumanBodyPosition(HumanBodyPosition string, ParenSpan opentracing.Span) bool {
	ChildSpan := opentracing.GlobalTracer().StartSpan(
		"body position",
		opentracing.ChildOf(ParenSpan.Context()),
		opentracing.Tag{Key: "body position", Value: HumanBodyPosition})
	defer ChildSpan.Finish()

	switch HumanBodyPosition {
	case "horizontally":
		return true //means that man fall down and can not get up
	case "vertically":
		return false //means that man stay on his legs
	}
	return false
}

// GetPedometer fix number of steps at this moment
func GetPedometer(Steps int, ParentSpan opentracing.Span) int {
	ChildSpan := opentracing.GlobalTracer().StartSpan(
		"get pedometer",
		opentracing.ChildOf(ParentSpan.Context()),
		opentracing.Tag{Key: "Num,er of steps", Value: Steps})
	defer ChildSpan.Finish()
	return Steps
}

// GetGroundImpact detects the fall of the body to the ground
func GetGroundImpact(GroundImpact bool, ParentSpan opentracing.Span) bool {
	ChildSpan := opentracing.GlobalTracer().StartSpan(
		"get ground impact",
		opentracing.ChildOf(ParentSpan.Context()),
		opentracing.Tag{Key: "impact sensor", Value: GroundImpact})
	defer ChildSpan.Finish()

	switch GroundImpact {
	case GroundImpact == true:
		return true
	case GroundImpact == false:
		return false
	}
	return false
}

// GetHeartRate - the sensor receives the heart rate and returns the probability of a heart attack
func GetHeartRate(HeartRate int, ParentSpan opentracing.Span) bool {
	ChildSpan := opentracing.GlobalTracer().StartSpan(
		"get heart rate",
		opentracing.ChildOf(ParentSpan.Context()),
		opentracing.Tag{Key: "heart rate", Value: HeartRate})
	defer ChildSpan.Finish()

	if HeartRate < 100 {
		return false
	} else if HeartRate >= 100 {
		return true
	} else {
		return false
	}
}

// GetBodyTemperature - the sensor receives temperature and returns the probability of a heart attack
func GetBodyTemperature(Temperature int, Parent opentracing.Span) bool {
	ChildSpan := opentracing.GlobalTracer().StartSpan(
		"get body temperature",
		opentracing.ChildOf(Parent.Context()),
		opentracing.Tag{Key: "body temperature", Value: Temperature})
	defer ChildSpan.Finish()

	if Temperature < 38 {
		return false
	} else if Temperature >= 38 {
		return true
	} else {
		return false
	}
}

// GetArterialPressure detects hypertension
func GetArterialPressure(Systolic int, Diastolic int, Parent opentracing.Span) bool {
	var Hypertension int

	if Systolic > 140 && Systolic < 160 && Diastolic > 90 && Diastolic < 100 {
		Hypertension = 1
	} else if Systolic >= 160 && Systolic < 180 && Diastolic > 100 && Diastolic < 110 {
		Hypertension = 2
	} else if Systolic >= 180 && Diastolic >= 110 {
		Hypertension = 3
	}

	ChildSpan := opentracing.GlobalTracer().StartSpan(
		"get arterial pressure",
		opentracing.ChildOf(Parent.Context()),
		opentracing.Tag{Key: "Systolic", Value: Systolic},
		opentracing.Tag{Key: "Diastolic", Value: Diastolic},
		opentracing.Tag{Key: "degree of hypertension", Value: Hypertension})
	defer ChildSpan.Finish()

	if Hypertension == 1 {
		return false
	} else if Hypertension == 2 {
		return false
	} else if Hypertension == 3 {
		return true
	} else {
		return false
	}
}

// GetDiabetes calculates the glucose rate and returns the probability of a critical situation
func GetDiabetes(glucose float64, Parent opentracing.Span) bool {
	ChildSpan := opentracing.GlobalTracer().StartSpan(
		"get diabetes",
		opentracing.ChildOf(Parent.Context()),
		opentracing.Tag{Key: "glucose", Value: glucose})
	defer ChildSpan.Finish()

	if glucose < 3.3 && glucose < 5.5 {
		return false
	} else if glucose >= 6.1 {
		return true
	}
	return false
}

func MessageToPatient(ParenSpan opentracing.Span) string {
	text := "Attention! \n\nYou are under the risk of a heart attack! \nDial your doctor's number immediately!" + ";" +
		"\n\n- Contact info: ..." + ";"

	ChildSpan := opentracing.GlobalTracer().StartSpan(
		"message to patient",
		opentracing.ChildOf(ParenSpan.Context()),
		opentracing.Tag{Key: "text message", Value: text})
	defer ChildSpan.Finish()

	return text
}

func MessageToDoctor(ParentSpan opentracing.Span) string {
	text := "Attention! \n\nThe patient is at high risk of heart attack! \nPlease call to patient" +
		"\n\n- Contact info: ..." + ";"

	ChildSpan := opentracing.GlobalTracer().StartSpan(
		"message to doctor",
		opentracing.ChildOf(ParentSpan.Context()),
		opentracing.Tag{Key: "message to doctor", Value: text})
	defer ChildSpan.Finish()

	return text
}

func MessageToEmergencyContact(ParentSpan opentracing.Span) string {
	text := "Attention! \n\nYour contact is at high risk of heart attack! \nContact him and clarify information about " +
		"the health of the contact! \nContact address:"
	ChildSpan := opentracing.GlobalTracer().StartSpan(
		"message to emergency contact",
		opentracing.ChildOf(ParentSpan.Context()),
		opentracing.Tag{Key: "message to contact", Value: text})
	defer ChildSpan.Finish()

	return text
}
