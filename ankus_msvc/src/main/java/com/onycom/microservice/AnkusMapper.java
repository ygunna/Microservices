package com.onycom.microservice;

import java.util.List;
import java.util.Map;

import org.apache.ibatis.annotations.Mapper;

@Mapper
public interface AnkusMapper {

	public String getNow();
	public List<Map<String,Object>> listMenu(Map<String,Object> cond);
	public int addMenu(Map<String,Object> data);
	public int updateMenu(Map<String,Object> data);
	public int deleteMenu(Map<String,Object> data);
	
	// 범용 crud
	public List<Map<String,Object>> listCommon(Map<String,Object> cond);
	public int listcountCommon(Map<String,Object> cond);
	public int addCommon(Map<String,Object> data);
	public int updateCommon(Map<String,Object> data);
	public int deleteCommon(Map<String,Object> data);

	// 창업플랫폼 api
	public List<Map<String,Object>> listDong(Map<String,Object> cond);
	public List<Map<String,Object>> listSangho(Map<String,Object> cond);
}
