data "external_schema" "gorm" {
  program = ["go", "run", "cmd/atlas/main.go"]
}

env "gorm" {
  src = data.external_schema.gorm.url
  dev = "docker://postgres/15"

  migration {
    dir = "file://migrations"
  }

  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}