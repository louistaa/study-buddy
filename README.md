# iStudy Buddy
## Project Description
iStudy Buddy is an application that will be primarily used by UW students, where students can use to find study buddies 
to study with for specific classes. Our application allows students to register for the classes that they are in, allowing 
them to find other students who are also looking for people to study with for specific classes. If a student cannot find 
the class that they'd like to register for, they have the ability to create the class that they'd like to find study buddies for.
If a student feels that they are an expert in a certain class, they can register themselves as an expert for the class and 
open themselves up to helping or tutoring other students. Once the quarter is over or if a student decides to drop a 
class, they have the ability to unregister themselves from a class.

Students will benefit from using this application because they will be able to easily find other students to work with 
(especially during online school!). 
## Technical Description

### Infrastructure

### Infrastructure Diagram

### Use cases and priority
| Priority    | User        | Description |
| ----------- | ----------- | ----------- |
| P0      | Student        | As a student, I want to view all available classes |
| P0   | Student        | As a student, I want to register for a class |
| P0   | Student        | As a student, I want to see who is in my class |
| P1   | Expert        | As an expert, I want to register as an expert for a class |
| P1   | Expert        | As an expert, I want  |

### API Design

`/students`:

`/students/`:
- POST
- GET

`/sessions/`:
- POST
- DELETE

`/classes`:
- POST
- DELETE

`/classes/{id}/people`:
- GET

`/classes/{id}/experts`:
- GET

`/register-class`:
- POST

`/register-expert`:
- POST


### Models


