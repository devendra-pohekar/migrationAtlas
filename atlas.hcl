data "external_schema" "beego" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "ariga.io/atlas-provider-beego",
    "load",
    "--path", "./models",
    "--dialect", "postgres", // | mysql | sqlite
  ]
}

env "beego" {
  src = data.external_schema.beego.url
  dev ="postgres://postgres:Dev@123@localhost:5432/atlas_migration"
  migration {
    dir = "file://migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}