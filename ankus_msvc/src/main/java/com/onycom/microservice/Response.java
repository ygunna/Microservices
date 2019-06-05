package com.onycom.microservice;

import java.util.List;

import lombok.Getter;
import lombok.Setter;

@Getter @Setter 
public class Response {

	private boolean success;
	private String message;
	private Object obj;
	private List<Object> list;
	private int total;
	private int limit;
}
