# People app

Build and run the app with following command in the root directory 
    ```
     docker-compose up --build 
    ```

Swagger docs will be served at http://127.0.0.1:2020/docs


#### Known issues
* Group ID inside user model is returned as sql.NullString object. Data transfer objects (input/output) are missing.
* Validation can be improved to indicate which parameter is invalid.
* User email is not required to be unique.
* More tests should be added (integration, unit).

...will be fixed in the next release :slightly_smiling_face:
