create table if not exists Student (
    ID int not null auto_increment primary key,
    Email varchar(255) not null unique, /* https://stackoverflow.com/questions/7717573/what-is-the-longest-possible-email-address  */
    PassHash binary(60) not null, /* https://stackoverflow.com/questions/5881169/what-column-type-length-should-i-use-for-storing-a-bcrypt-hashed-password-in-a-d */
    UserName varchar(255) not null unique, 
    FirstName varchar(128) not null,
    LastName varchar(128) not null,
    PhotoURL varchar(2083) not null,
    Major varchar(300) not null,
);

create table if not exists SignIn (
    session_pk int not null auto_increment primary key,
    ID int not null,
    SigninDateTime datetime not null,
    IPAddress varchar(60) not null UNIQUE
);

create table if not exists Courses (
    ID int not null auto_increment primary key,
    CourseName varchar(128) not null,
    CourseDept varchar(128) not null,
    ProfessorID int not null
)

create table if not exists Professor (
  ID int not null auto_increment primary key,
  ProfessorFirstName varchar(128) not null,
  ProfessorLastName varchar(128) not null,
)

create table if not exists StudentCourse (
  ID int not null auto_increment primary key,
  StudentID int not null,
  CourseID int not null, 
  isCurrentCourse boolean not null,
  isExpert boolean not null   
)

create table if not exists Ratings (
    ID int not null auto_increment primary key,
    RatingDifficulty int not null,
    RatingEnjoyment int not null,
    RatingAvgTimeConsumptionPerWeek int not null,
    RatingInstructorEngagement int not null
)

create table if not exists RatingCourse (
    ID int not null auto_increment primary key,
    RatingID int not null,
    CourseID int not null
)