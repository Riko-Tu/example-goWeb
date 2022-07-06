
-- //查询"01"课程比"02"课程成绩高的学生的信息及课程分数
select s.*,a.s_score as chinses,b.s_score as math  from Score a ,Score b ,Student s
where a.c_id='01' and b.c_id ='02' and a.s_id = b.s_id and a.s_score >b.s_score and a.s_id = s.s_id;



-- 2、查询"01"课程比"02"课程成绩低的学生的信息及课程分数
select S4.*, s.s_score,s2.s_score as s02
from
    (select * from Score union all select '07' as s_id, '01'as c_id, '0' as s_score union all  select  '06' as s_id,'02'as c_id,'0' as s_score) s ,
    (select * from Score union all select '07' as s_id, '01' as c_id, '0' as s_score)  s2 , Student S4
where s.c_id ='01' and  s2.c_id ='02' and s.s_id =s2.s_id and s.s_score < s2.s_score and s.s_id = S4.s_id;

-- 3、查询平均成绩大于等于60分的同学的学生编号和学生姓名和平均成绩
SELECT S2.s_id,S.s_name, ROUND(avg(S2.s_score),1) as avg FROM Student S ,Score S2
WHERE S2.s_id=S.s_id GROUP BY S2.s_id having avg >=60;

--
-- # 4、查询平均成绩小于60分的同学的学生编号和学生姓名和平均成绩
-- # 		-- (包括有成绩的和无成绩的)


select s2.s_id,s.s_name, round(avg(s2.s_score),2) as avg  from Student s , Score s2
where s.s_id =s2.s_id group by s2.s_id , s.s_name having avg<=60
union select s_id ,s_name,0 as avg from Student
      where s_id not in (select distinct  s_id from Score  ) ;


-- 5、查询所有同学的学生编号、学生姓名、选课总数、所有课程的总成绩

select a.s_id,b.s_name, count(a.c_id) as course ,sum(a.s_score) as total
from Score a join Student b on a.s_id=b.s_id group by a.s_id ,b.s_name


-- 7、查询学过"张三"老师授课的同学的信息


select * from Student
where s_id in  (select s_id from Score where  c_id =
( select c_id from  Course where t_id =(select t_id from Teacher where t_name = '张三')) group by s_id)
-- 8、查询没学过"张三"老师授课的同学的信息

select *
from Student where s_id not in( select s_id
                            from Score where  c_id =
                                                  ( select c_id from  Course where t_id =(select t_id from Teacher where t_name = '张三')) group by s_id);


-- 9、查询学过编号为"01"并且也学过编号为"02"的课程的同学的信息


select b.*
from Score a  ,Score c , Student b where a.s_id =b.s_id = c.s_id and  a.c_id= '02' and c.c_id = '01';

-- 10、查询学过编号为"01"但是没有学过编号为"02"的课程的同学的信息

select * from Student where s_id in (select s_id from Score   where c_id ='01' and  s_id not in   (select s_id  from Score where c_id ='02')) ;
select a.* from Score b ,Student a where  b.c_id in ('01') and b.s_id not in (select s_id from Score where c_id='02') and a.s_id = b.s_id ;

-- 11、查询没有学全所有课程的同学的信息

select *
from Student  where s_id not in (select s_id from Score group by s_id having count(c_id)=(select count(*) from Course));


-- 12、查询至少有一门课与学号为"01"的同学所学相同的同学的信息

select * from Student where s_id in (select distinct s_id
                                     from Score where c_id in (select c_id
                                                               from Score where  s_id = '01'));



-- 13、查询和"01"号的同学学习的课程完全相同的其他同学的信息



select * from Student where s_id in (select s_id
from Score where c_id in (select c_id from Score where s_id = '01' ) and  s_id != '01'
group by s_id having  count(s_id)=(select count(c_id) from Score where s_id = '01') ) ;




-- 14、查询没学过"张三"老师讲授的任一门课程的学生姓名



select s_name from Student where s_id not  in (select distinct  s_id from Score where c_id = (select c_id from Course where t_id =
(select t_id from Teacher where t_name ='张三')));




-- 15、查询两门及其以上不及格课程的同学的学号，姓名及其平均成绩

select b.s_id,b.s_name, round(avg(a.s_score),2) avgScore from Score a, Student b
where a.s_score < 60 and a.s_id = b.s_id group by  a.s_id having  count(a.c_id)>=2 ;





