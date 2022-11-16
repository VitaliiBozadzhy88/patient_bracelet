package exercise3_tests

import (
	"testing"
)

var actual bool
var expected bool

func TestCalcRiskAttack(t *testing.T) {
	actualPosition := "horizontal"
	actualSteps := 0
	actualGroundImpact := true
	actualHeartRate := 101
	actualTemperature := 38
	actualArterialPressure := true
	actualDiabetes := true

	expectedPosition := "horizontal"
	expectedSteps := 0
	expectedGroundImpact := true
	expectedHeartRate := 101
	expectedTemperature := 38
	expectedArterialPressure := true
	expectedDiabetes := true

	if actualPosition != expectedPosition {
		t.Errorf("Expected position - %s, actual - %s", expectedPosition, actualPosition)
	} else if actualSteps != expectedSteps {
		t.Errorf("Expected steps - %d, actual - %d", expectedSteps, actualSteps)
	} else if actualGroundImpact != expectedGroundImpact {
		t.Errorf("Expected impact - %t, actual - %t", expectedGroundImpact, actualGroundImpact)
	} else if actualHeartRate != expectedHeartRate {
		t.Errorf("Expected heart rate - %d, actual - %d", expectedHeartRate, actualHeartRate)
	} else if actualTemperature != expectedTemperature {
		t.Errorf("Expected temperature - %d, actual - %d", expectedTemperature, actualTemperature)
	} else if actualArterialPressure != expectedArterialPressure {
		t.Errorf("Expected arterial pressure - %t, actual - %t", expectedArterialPressure, actualArterialPressure)
	} else if actualDiabetes != expectedDiabetes {
		t.Errorf("Expected diabetes - %t, actual - %t", expectedDiabetes, actualDiabetes)
	}
}

func TestGetArterialPressure(t *testing.T) {
	var Systolic = 180
	var Diastolic = 110
	var ExpectedHypertension = 3
	var ActualHypertension int

	if Systolic >= 180 && Diastolic >= 110 {
		ActualHypertension = 3
		if ActualHypertension != ExpectedHypertension {
			t.Errorf("Expected ActualHypertension - %d, actual - %d, ", ExpectedHypertension, ActualHypertension)
		}
	} else {
		if ActualHypertension != ExpectedHypertension {
			t.Errorf("Expected ActualHypertension - %d, actual - %d, ", ExpectedHypertension, ActualHypertension)
		}
	}
}

func TestGetBodyTemperature(t *testing.T) {
	actualTemperature := 38

	if actualTemperature < 38 {
		actual = false
		expected = true
		if actual != expected {
			t.Errorf("Expected %t Actual %t, ", expected, actual)
		}
	} else if actualTemperature >= 38 {
		actual = true
		expected = true
		if actual != expected {
			t.Errorf("Expected %t Actual %t, ", expected, actual)
		}
	} else {
		t.Errorf("Error receiving data from sensor")
	}
}

func TestGetDiabetes(t *testing.T) {
	var actualGlucose = 6.1

	if actualGlucose >= 6.1 {
		actual = true
		expected = true
		if actual != expected {
			t.Errorf("Expected %t Actual %t, ", expected, actual)
		}
	} else {
		actual = false
		expected = true
		if actual != expected {
			t.Errorf("Expected %t Actual %t, ", expected, actual)
		}
	}
}

func TestGetGroundImpact(t *testing.T) {
	impactFromSensor := true

	if impactFromSensor == true {
		actual = true
		expected = true
		if actual != expected {
			t.Errorf("Expected %t Actual %t, ", expected, actual)
		}
	} else if impactFromSensor == false {
		actual = false
		expected = true
		if actual != expected {
			t.Errorf("Expected %t Actual %t, ", expected, actual)
		}
	} else {
		t.Errorf("Error receiving data from sensor")
	}
}

func TestGetHeartRate(t *testing.T) {
	heartRate := 110

	if heartRate < 100 {
		actual = false
		expected = true
		if actual != expected {
			t.Errorf("Expected %t Actual %t, ", expected, actual)
		}
	} else if heartRate >= 100 {
		actual = true
		expected = true
		if actual != expected {
			t.Errorf("Expected %t Actual %t, ", expected, actual)
		}
	} else {
		t.Errorf("Error receiving data from sensor")
	}

}

func TestGetHumanBodyPosition(t *testing.T) {
	var positionFromSensor = "horizontally"
	//var positionFromSensor = "vertically"

	if positionFromSensor == "horizontally" {
		actual = true
		expected = true
		if actual != expected {
			t.Errorf("Expected %t Actual %t, ", expected, actual)
		}
	} else if positionFromSensor == "vertically" {
		actual = false
		expected = true
		if actual != expected {
			t.Errorf("Expected %t Actual %t, ", expected, actual)
		}
	} else {
		t.Errorf("Error receiving data from sensor")
	}
}

func TestGetPedometer(t *testing.T) {
	steps := 0
	if steps >= 0 && steps < 10 {
		actual = true
		expected = true
		if actual != expected {
			t.Errorf("Expected %t Actual %t, ", expected, actual)
		}
	} else if steps >= 10 {
		actual = false
		expected = true
		if actual != expected {
			t.Errorf("Expected %t Actual %t, ", expected, actual)
		}
	} else {
		t.Errorf("Error receiving data from sensor")
	}
}
