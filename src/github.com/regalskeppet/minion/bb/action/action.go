package action

import "errors"

func GetAction(action string) (func(args []string) error, error) {
  switch action {
  case "help":
    return Help, nil
  case "apply":
    return Apply, nil
  }
  return nil, errors.New("Unrecognized action.")
}
