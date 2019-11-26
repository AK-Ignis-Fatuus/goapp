<<<<<<< HEAD
# Go_todo_app
Repo for my app on Golang

Routs:
/tasks/ - [GET] get all tasks
/tasks/create?taskID={}&taskName={}&Todo={} - [POST] create new task
/tasks/{taskID} [DELETE] - delete task by id
/tasks/deleteall [DELETE] - delete all tasks

DB:

CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    taskid character varying(50) NOT NULL,
    taskname character varying(255) NOT NULL,
    todo character varying(255) NOT NULL
);




=======
# goapp
>>>>>>> 598cd2468ccdee709d6fce82148d7708c8b27611
