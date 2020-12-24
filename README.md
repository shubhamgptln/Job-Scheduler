# Job-Scheduler
Job scheduling framework for distributed systems

Introduction

Tha main aim is to design and develop a distributed task scheduler where we can spawn any nymber of task and wait for the execution to complete. The main aim is to 
design it to perform big operation like database migration or reindexing Elastic Search index which may contain billions of record. The operation should be performed with few hours or a couple of days depending on the priority of task. To finsh the task within stipulated time we need to deploy N number of machines where N may vary depending on amount of task and compute required to execute each task
