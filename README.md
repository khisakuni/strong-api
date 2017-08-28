# Strong API

### Endpoints

##### GET `/api/v1/workouts`

Fetches list of `workout`s


##### GET `/api/v1/workouts/:id`

Fetches individual workout


##### POST `/api/v1/workouts`

Creates workout.

_Request params:_

```
{
  "name": "Example workout",
  "description": "A description of the workout",
  "instagramId": "shortcode"
}
```

The `instagramId` is optional, and is the _short code_, found in the
url of an instagram post. (https://www.instagram.com/developer/embedding/)


##### PUT `/api/v1/workouts/:id`

Updates `workout` record.

_Reuqest params:_

```
{
  "name": "Example workout",
  "description": "A description of the workout",
  "instagramId": "shortcode"
}
```


### Deployment

Deploy to heroku using git.

##### Migrations

This project uses [sql-migrate](https://github.com/rubenv/sql-migrate) to manage migrations.
To create a migration, run:
```
$ sql-migrate new YourMigrationHere
```
then edit the generated sql file.
Apply the migration locally by running:
```
$ sql-migrate up
```
To undo the migraion, run:
```
$ sql-migrate down
```

Apply migrations on heroku by running:
```
$ heroku run migrations -a <heroku-instance-name>




