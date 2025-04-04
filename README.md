Soroush Azhari
soroushazhari566@gmail.com

This is my first project both on Go and GitHub.
The target is to implement a task manager. Multi-Users can Signup / SignIn and make tasks and task-categories related to each other
My target is to implement server / client by TCP connection and while client works with CLI, server receives command by json-serialized TCP packets
It's now very simple without any database

Business logics:

Entities: 
    Category:
        ID
        Title
        Progress

        Create()
        List() // list categories associated to a user
        Edit()

    Task:
        ID
        Title
        IsDone
        Category
        DeadLine

        Create()
        List() // list tasks associated to a user / categories / date
        MarkDone()
        Edit()

    User:
        ID
        Email
        Password

        SignUp()
        SignIn()