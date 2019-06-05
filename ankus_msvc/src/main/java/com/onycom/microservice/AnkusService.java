package com.onycom.microservice;

import java.util.List;
import java.util.Map;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class AnkusService {

	@Autowired
	AnkusMapper demomapper;
	
	public String getNow()
	{
		return demomapper.getNow();
	}

	public List<Map<String,Object>> listMenu(Map<String,Object> cond)
	{
		return demomapper.listMenu(cond);
	}

	public int addMenu(Map<String,Object> data)
	{
		return demomapper.addMenu(data);
	}

	public int updateMenu(Map<String,Object> data)
	{
		return demomapper.updateMenu(data);
	}

	public int deleteMenu(Map<String,Object> data)
	{
		return demomapper.deleteMenu(data);
	}

	public List<Map<String,Object>> listCommon(Map<String,Object> cond)
	{
		return demomapper.listCommon(cond);
	}

	public int listcountCommon(Map<String,Object> cond)
	{
		return demomapper.listcountCommon(cond);
	}
	
	public int addCommon(Map<String,Object> data)
	{
		return demomapper.addCommon(data);
	}

	public int updateCommon(Map<String,Object> data)
	{
		return demomapper.updateCommon(data);
	}

	public int deleteCommon(Map<String,Object> data)
	{
		return demomapper.deleteCommon(data);
	}

	// 창업플랫폼 api
	public List<Map<String,Object>> listDong(Map<String,Object> cond)
	{
		return demomapper.listDong(cond);
	}
	
	public List<Map<String,Object>> listSangho(Map<String,Object> cond)
	{
		return demomapper.listSangho(cond);
	}
	
}
