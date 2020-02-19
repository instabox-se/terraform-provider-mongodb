provider "mongodb" {
  // Please note that you need a user with userAdmin role or similar to
  // be able to use this provider. It is not recommended to manage this user
  // through Terraform :)
  connection_string = "mongodb://localhost:27017"
}

resource "mongodb_user" "user1" {
  username = "user1"
  db       = "admin"
  name     = "User Userson"
  password = "<make sure this is not part of state...>"

  role {
    role = mongodb_role.ib_developer.name
    db   = "admin"
  }
}

resource "mongodb_role" "ib_developer" {
  name = "ib-developer"
  db   = "admin"

  role {
    role = "read"
    db   = "production"
  }
  role {
    role = "readWrite"
    db   = "development"
  }
  role {
    role = "readWrite"
    db   = "sandbox"
  }

  privilege {
    cluster = true
    actions = ["listDatabases"]
  }

  privilege {
    db = "admin"
    collection = "*"
    actions = ["find"]
  }
}
