# crud-go-server
Go server using go-chi and dependency injection

## API
Document for all CRUD APIs 

### Get All Student list
Used to collect full info about students

**URL**: `/student`

**Method**: `GET`

Success Response example:
```json
[
    {
        "studentID": "20060101",
        "surname": "Dickens",
        "forename": "Charles",
        "scores": [
            {
                "module_name": "Databases",
                "module_code": "CM0001",
                "mark": 80
            }
        ]
    }
]
```

### Get Student by studentID
Retrieve full student info by studentID

**URL**: `/student/{studentID}`

**Method**: `GET`

Success Response example:
```json
[
    {
        "studentID": "20060101",
        "surname": "Dickens",
        "forename": "Charles",
        "scores": [
            {
                "module_name": "Databases",
                "module_code": "CM0001",
                "mark": 80
            }
        ]
    }
]
```

#### Error Response
**Condition**: Not exist studentID 

**Code**: `404 NOT FOUND`

**Content example**
```json
{
    "status": "Resource not found."
}
```

### Add new student
Create a new Student if studentID is not used. No 2 students use the same StudentID

**URL**: `student/create`

**Method**: `POST`

Provide info of Student to be created
```json
{
    "studentID": "111",
    "surname": "Pham",
    "forename": "Nhan"
}
```

Success Response:
```json
{
    "studentID": "111",
    "surname": "Pham",
    "forename": "Nhan"
}
```

#### Error Response
**Condition**: Already exist studentID 

**Code**: `400 BAD REQUEST`

**Content example**
```json
{
    "status": "Invalid Request",
    "error": "Error 1062: Duplicate entry '111' for key 'students.PRIMARY'"
}
```

### Update student
Update student info by studentID

**URL**: `student/{studentID}`

**Method**: `PUT`

Provide info of Student to be updated
```json
{
    "surname": "Trong",
    "forename": "Nhan"
}
```

Success Response:
```json
{
    "studentID": "111",
    "surname": "Trong",
    "forename": "Nhan"
}
```

### Delete Student
Delete student info

**URL**: `student/delete`

**Method**: `DELETE`

Provide info of Student to be created
```json
{
    "id": "111"
}
```

Success Response:
```json
null
```

### Get Mark
Used to collect full info about mark by student or module

**URL**: `/mark`

**Method**: `GET`

**Optional query**: 
- `student_no`: query by studentID
- `module_code`: query by moduleID

Success Response example:
```json
[
    {
        "student_no": "20060101",
        "module_code": "CM0001",
        "mark": 80
    },
    {
        "student_no": "20060101",
        "module_code": "CM0002",
        "mark": 65
    },
    {
        "student_no": "20060101",
        "module_code": "CM0003",
        "mark": 50
    },
    {
        "student_no": "20060101",
        "module_code": "CM0004",
        "mark": 87
    }
]
```

### Add new mark
Add new mark. If student_no or module_code are not existed in tables Students/Modules return error

**URL**: `mark/create`

**Method**: `POST`

Provide info of Mark to be created
```json
{
    "student_no": "111",
    "module_code": "CM0005",
    "mark": 87
}
```

Success Response:
```json
{
    "student_no": "111",
    "module_code": "CM0005",
    "mark": 87
}
```

#### Error Response
**Condition**: Not exist student_no or module_code

**Code**: `400 BAD REQUEST`

**Content example**
```json
{
    "status": "Invalid Request",
    "error": msg
}
```

### Update Mark
Update student's mark by student_no and module_code. If student_no/module_code is not existed then return `200 OK` but with no change on database

**URL**: `student/`

**Method**: `PUT`

Provide info of Mark to be updated
```json
{
    "student_no": "111",
    "module_code": "CM0005",
    "mark": 87
}
```

Success Response:
```json
{
    "student_no": "111",
    "module_code": "CM0005",
    "mark": 87
}
```

### Delete Mark
Delete mark info

**URL**: `mark/delete`

**Method**: `DELETE`

Provide info of Mark to be created
```json
{
    "student_no": "111",
    "module_code": "CM0005",
}
```

Success Response:
```json
null
```

### Get All Module list
Used to collect full info about modules

**URL**: `/module`

**Method**: `GET`

Success Response example:
```json
[
    {
        "module_code": "CM0001",
        "module_name": "Databases"
    }
]
```

### Get Module by module_code
Retrieve full module info by module_code

**URL**: `/module/{module_code}`

**Method**: `GET`

Success Response example:
```json
{
    "module_code": "CM0001",
    "module_name": "Databases"
}
```

#### Error Response
**Condition**: Not exist module_code 

**Code**: `404 NOT FOUND`

**Content example**
```json
{
    "status": "Resource not found."
}
```

### Add new module
Create a new Module if module_code is not used. No 2 modules use the same module_code

**URL**: `module/create`

**Method**: `POST`

Provide info of Student to be created
```json
{
    "module_code": "CSE007",
    "module_name": "Network",
}
```

Success Response:
```json
{
    "module_code": "CSE007",
    "module_name": "Network"
}
```

#### Error Response
**Condition**: Already exist studentID 

**Code**: `400 BAD REQUEST`

**Content example**
```json
{
    "status": "Invalid Request",
    "error": msg
}
```

### Update module
Update module info by module_code

**URL**: `module/{module_code}`

**Method**: `PUT`

Provide info of Module to be updated with example: `/module/CSE007`
```json
{
    "module_name": "Security"
}
```

Success Response:
```json
{
    "module_code": "CSE007",
    "module_name": "Security"
}
```

### Delete Module
Delete module info

**URL**: `module/delete`

**Method**: `DELETE`

Provide info of Module to be deleted
```json
{
    "module": "CSE007"
}
```

Success Response:
```json
null
```