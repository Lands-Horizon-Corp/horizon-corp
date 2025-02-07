feature/api-service-layer - switch if modifying/creating api communication related files & types - src/hooks/api-hooks - src/server/\*

feature/forms
|\_\_\_ Rely on 71-feature/api-service-layer - switch if modifying/creating forms or creating zod validation for forms - src/components/forms/_ - src/validation/_

feature/tables
|\_\_\_ Rely on 71-feature/api-service-layer - switch if modifying/creating tables related files/components - src/components/tables/_ - src/components/table/_

feature/owner
|\_\_\_ Rely on feature/forms, 71-feature/api-service-layer, 72-feature/tables - switch if doing owner related features

feature/owner-accounting
|\_\_\_ Rely on feature/forms, 71-feature/api-service-layer, 72-feature/tables - owner accounting feature

feature/components
|\_\_\_ ... - switch here if making changes to general components especially src/components/ui
