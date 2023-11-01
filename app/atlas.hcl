data "external_schema" "gorm" {
  program = ["go", "run", "cmd/atlas/main.go"]
}

env "gorm" {
  src = data.external_schema.gorm.url
  dev = "docker://mysql/8/schema"

  migration {
    dir = "file://migrations"
  }

  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}