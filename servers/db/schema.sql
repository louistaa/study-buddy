create table if not exists Students (
    ID int not null auto_increment primary key,
    Email varchar(255) not null unique, /* https://stackoverflow.com/questions/7717573/what-is-the-longest-possible-email-address  */
    PassHash binary(60) not null, /* https://stackoverflow.com/questions/5881169/what-column-type-length-should-i-use-for-storing-a-bcrypt-hashed-password-in-a-d */
    UserName varchar(255) not null unique, 
    FirstName varchar(128) not null,
    LastName varchar(128) not null,
    PhotoURL varchar(2083) not null,
    Major varchar(300) not null
);

create table if not exists SignInLog (
    ID int not null auto_increment primary key,
    StudentID int not null,
    Time DateTime not null,
    ClientIP varchar(255) not null
);

create table if not exists Courses (
    ID int not null auto_increment primary key,
    Name varchar(128) not null,
    Department varchar(128) not null,
    ProfessorID int,
    FOREIGN KEY (ProfessorID) REFERENCES Professors(ID)
);

create table if not exists Professors (
  ID int not null auto_increment primary key,
  FirstName varchar(128) not null,
  LastName varchar(128) not null,
);

create table if not exists StudentCourse (
  ID int not null auto_increment primary key,
  StudentID int,
  CourseID int,
  StartDate DATE not null,
  EndDate DATE not null,
  FOREIGN KEY (StudentID) REFERENCES Students(ID),
  FOREIGN KEY (CourseID) REFERENCES Courses(ID)
);

create table if not exists CourseExpert (
    ID int not null auto_increment primary key,
    CourseID int,
    ExpertID int,
    FOREIGN KEY (CourseID) REFERENCES Courses(ID),
    FOREIGN KEY (ExpertID) REFERENCES Students(ID)
);

create table if not exists Ratings (
    ID int not null auto_increment primary key,
    Difficulty int not null,
    Enjoyment int not null,
    AvgTimeConsumptionPerWeek int not null,
    InstructorEngagement int not null
);

create table if not exists RatingCourse (
    ID int not null auto_increment primary key,
    RatingID int not null,
    CourseID int not null
);