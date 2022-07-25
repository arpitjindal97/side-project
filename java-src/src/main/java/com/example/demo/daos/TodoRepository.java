package com.example.demo.daos;

import com.example.demo.entities.Todo;
import org.springframework.data.repository.CrudRepository;
import org.springframework.data.rest.core.annotation.RepositoryRestResource;

@RepositoryRestResource
public interface TodoRepository extends CrudRepository<Todo,String> {

}