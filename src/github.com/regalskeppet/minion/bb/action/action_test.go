package action

import "testing"

func TestActionsExist(t *testing.T) {
  actions := []string{
    "help",
    "apply",
  }
  for _, a := range actions {
    testGetActionExist(a, true, t)
  }
}

func TestActionsUnrecognized(t *testing.T) {
  actions := []string{
    "unrecognized",
  }
  for _, a := range actions {
    testGetActionExist(a, false, t)
  }
}

func testGetActionExist(action string, exist bool, t *testing.T) {
  _, err := GetAction(action)
  if exist && err != nil {
    t.Error(action, err)
  } else if !exist && err == nil {
    t.Error(action, "Found action that should not exist.")
  }
}
