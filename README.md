# Go_todo_app
Repo for my app on Golang

Routs:<br>
/tasks/ - [GET] get all tasks <br>
/tasks/create?taskID={}&taskName={}&Todo={} - [POST] create new task <br>
/tasks/{taskID} [DELETE] - delete task by id <br>
/tasks/deleteall [DELETE] - delete all tasks <br>

DB: <br>

CREATE TABLE tasks ( <br>
    id SERIAL PRIMARY KEY, <br>
    taskid character varying(50) NOT NULL, <br>
    taskname character varying(255) NOT NULL, <br>
    todo character varying(255) NOT NULL <br>
); <br>




=======
# goapp
