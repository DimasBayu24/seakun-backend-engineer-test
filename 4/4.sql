DROP TABLE IF EXISTS admin_test;
DROP TABLE IF EXISTS teacher_class_test;
DROP TABLE IF EXISTS teacher_test;
DROP TABLE IF EXISTS class_test;

CREATE TABLE teacher_test(
   teacher_id INT GENERATED ALWAYS AS IDENTITY,
   teacher_name VARCHAR(255) NOT NULL,
   PRIMARY KEY(teacher_id)
);

CREATE TABLE class_test(
   class_id INT GENERATED ALWAYS AS IDENTITY,
   class_name VARCHAR(255) NOT NULL,
   PRIMARY KEY(class_id)
);

CREATE TABLE teacher_class_test(
    teacher_id INT,
    class_id INT,
   CONSTRAINT fk_teacher
      FOREIGN KEY(teacher_id) 
	  REFERENCES teacher_test(teacher_id)
	  ON DELETE SET NULL,
      CONSTRAINT fk_class
      FOREIGN KEY(class_id) 
	  REFERENCES class_test(class_id)
	  ON DELETE SET NULL
);

CREATE TABLE admin_test(
   admin_test_id INT GENERATED ALWAYS AS IDENTITY,
    teacher_id INT,
   PRIMARY KEY(admin_test_id),
   CONSTRAINT fk_teacher
      FOREIGN KEY(teacher_id) 
	  REFERENCES teacher_test(teacher_id)
	  ON DELETE SET NULL
);

INSERT INTO teacher_test(teacher_name)
VALUES('Camat Matamata'),
      ('Ajengskuy');

INSERT INTO class_test(class_name)
VALUES('1A'),
      ('2A');	

INSERT INTO teacher_class_test(teacher_id, class_id)
VALUES(1,1),
      (2,2);	   
	   
INSERT INTO admin_test(teacher_id)
VALUES(1),
      (2);