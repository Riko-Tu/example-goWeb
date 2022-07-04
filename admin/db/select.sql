
-- //查询"01"课程比"02"课程成绩高的学生的信息及课程分数
-- 方式1
select a.* ,b.s_score as 01_score,c.s_score as 02_score from
    Student a
        join Score b on a.s_id=b.s_id and b.c_id='01'
        left join Score c on a.s_id=c.s_id and c.c_id='02' or c.c_id = NULL where b.s_score>c.s_score
--  方式2
(select s.*,a.s_score as chinses,b.s_score as math  from Score a ,Score b ,Student s
where a.c_id='01' and b.c_id ='02' and a.s_id = b.s_id and a.s_score >b.s_score and a.s_id = s.s_id)



-- 2、查询"01"课程比"02"课程成绩低的学生的信息及课程分数
select S4.*, s.s_score,s2.s_score as s02
from
    (select * from Score union all select '07' as s_id, '01'as c_id, '0' as s_score union all  select  '06' as s_id,'02'as c_id,'0' as s_score) s ,
    (select * from Score union all select '07' as s_id, '01' as c_id, '0' as s_score)  s2 , Student S4
where s.c_id ='01' and  s2.c_id ='02' and s.s_id =s2.s_id and s.s_score < s2.s_score and s.s_id = S4.s_id;
