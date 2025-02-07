feature/api-service-layer
     - switch if modifying/creating api communication related files & types
     - src/hooks/api-hooks
     - src/server/*

feature/forms
|___ Rely on 71-feature/api-service-layer
    - switch if modifying/creating forms or creating zod validation for forms
    - src/components/forms/*
    - src/validation/*

feature/tables
|___ Rely on 71-feature/api-service-layer
     - switch if modifying/creating tables related files/components
     - src/components/tables/*
     - src/components/table/*

feature/owner
|___ Rely on feature/forms, 71-feature/api-service-layer, 72-feature/tables
    - switch if doing owner related features

feature/owner-accounting
|___ Rely on feature/forms, 71-feature/api-service-layer, 72-feature/tables
    - owner accounting feature

feature/components
|___ ...
    - switch here if making changes to general components especially src/components/ui