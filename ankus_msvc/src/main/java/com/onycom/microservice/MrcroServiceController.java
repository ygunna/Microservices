package com.onycom.microservice;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.cloud.context.config.annotation.RefreshScope;
import org.springframework.web.bind.annotation.CrossOrigin;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.ResponseBody;
import org.springframework.web.bind.annotation.RestController;


@CrossOrigin("*")
@RestController
@RequestMapping("/api")
@RefreshScope
public class MrcroServiceController {

	private static final Logger logger = LoggerFactory.getLogger(MrcroServiceController.class);

	@Autowired
	AnkusService demoservice;

	/*
	@RequestMapping(value="/get", method = RequestMethod.POST)
	public List<Data> search(Map model) {
		System.out.println("request /serarc/get");

		Data data = new Data();
		data.setId("1");
		data.setName("name");

		List<Data> list = new ArrayList<Data>();
		list.add(data);

		return list;
	}

	@RequestMapping(value="/getTwo", method = RequestMethod.POST)
	public List<Data> searchTwo(Map model) {
		System.out.println("request /serarc/getTwo");

		Data data = new Data();
		data.setId("1");
		data.setName("name");

		List<Data> list = new ArrayList<Data>();
		list.add(data);

		data = new Data();
		data.setId("2");
		data.setName("name");

		list.add(data);

		return list;
	}

	@RequestMapping(value="/getThree", method = RequestMethod.POST)
	public List<Data> searchThree(Map model) {
		System.out.println("request /serarc/getThree");

		Data data = new Data();
		data.setId("1");
		data.setName("name");

		List<Data> list = new ArrayList<Data>();
		list.add(data);

		data = new Data();
		data.setId("2");
		data.setName("name");

		list.add(data);

		data = new Data();
		data.setId("3");
		data.setName("name");

		list.add(data);

		return list;
	}
	*/

	/*
    @RequestMapping(value="/{table}/list", method=RequestMethod.POST)
    // , produces={"application/json","application/xml"}, consumes="text/html")
    @ResponseBody
    public Response list(@PathVariable("table") String table, @RequestParam Map<String,Object> cond) {
    	
    	Map<String,Object> searchcond = new HashMap<String,Object>();
    	searchcond.put("table", table);
    	
    	List<String> conds = new ArrayList<String>();
    	List<Object> cond_vals = new ArrayList<Object>();
    	
    	String vals = (String)cond.get("conds");
    	
    	List<String> fldlst = new ArrayList<String>();
    	if(vals!=null) for(String i:vals.split(",")) if(!i.trim().isEmpty()) fldlst.add(i.trim());
    	
//    	searchcond.put("conds", fldlst);
    	
    	vals = (String)cond.get("orderby");
    	
    	List<String> ordlst = new ArrayList<String>();
    	if(vals!=null) for(String i:vals.split(",")) if(!i.trim().isEmpty()) ordlst.add(i.trim());
    	
    	searchcond.put("orderby", ordlst);
    	
    	for(String key:fldlst)
    	{
    		String value = (String)cond.get(key);

        	List<String> values = new ArrayList<String>();
        	
        	if(value!=null) for(String i:value.split(",")) if(!i.trim().isEmpty()) values.add(i.trim());

			conds.add(key);
			cond_vals.add(values);
    	}
    	
    	String groupby = (String)cond.get("groupby");
    	
    	if(groupby!=null) searchcond.put("groupby", groupby.split(","));
    	
    	searchcond.put("conds", conds);
    	searchcond.put("cond_vals", cond_vals);
    	
    	searchcond.put("recordstartindex", cond.get("recordstartindex"));
    	searchcond.put("pagesize", cond.get("pagesize"));
    	
    	int total = -1;
    	
    	if(cond.get("total")==null || ((String)cond.get("total"))=="")
    	{
    		total = demoservice.listcountCommon(searchcond);
    	}
    	else
    	{
    		total = Integer.parseInt(((String)cond.get("total")));
    	}
    	
    	List<Map<String, Object>> lst = demoservice.listCommon(searchcond);
    	
    	System.out.printf("cond==>%s/%s, lst=%s\n", cond.toString(), searchcond.toString(), lst.size());
    	
    	Response resp = new Response();
    	
    	resp.setTotal(total);
    	
    	resp.setSuccess(true);
    	resp.setList((List) lst);
    	
        return resp;
    }

    @RequestMapping(value="/{table}/listcount", method=RequestMethod.POST)
    // , produces={"application/json","application/xml"}, consumes="text/html")
    @ResponseBody
    public Response listcount(@PathVariable("table") String table, @RequestParam Map<String,Object> cond) {

    	Map<String,Object> searchcond = new HashMap<String,Object>();
    	searchcond.put("table", table);
    	
    	List<String> conds = new ArrayList<String>();
    	List<Object> cond_vals = new ArrayList<Object>();
    	
    	String vals = (String)cond.get("conds");
    	
    	List<String> fldlst = new ArrayList<String>();
    	if(vals!=null) for(String i:vals.split(",")) if(!i.trim().isEmpty()) fldlst.add(i.trim());
    	
//    	searchcond.put("conds", fldlst);
    	
    	for(String key:fldlst)
    	{
    		String value = (String)cond.get(key);

        	List<String> values = new ArrayList<String>();
        	
        	if(value!=null) for(String i:value.split(",")) if(!i.trim().isEmpty()) values.add(i.trim());

			conds.add(key);
			cond_vals.add(values);
    	}
    	
    	searchcond.put("conds", conds);
    	searchcond.put("cond_vals", cond_vals);
    	
    	System.out.printf("cond==>%s\n", searchcond.toString());
    	
    	Response resp = new Response();
    	
    	resp.setSuccess(true);
    	resp.setTotal(demoservice.listcountCommon(searchcond));
    	
        return resp;
    }
    
    @RequestMapping("/{table}/add")
    @ResponseBody
    public Response add(@PathVariable("table") String table, @RequestParam Map<String,Object> data) {
    	
    	Response resp = new Response();
    	
    	data.put("table", table);
//    	data.put("id", generateId("MENU_"));
//    	columns, values  처리...
    	
    	String[] columns = ((String)data.get("columns")).split(",");

    	data.put("columns", columns);
    	
    	if(data.get("values")!=null)
    		data.put("values", ((String)data.get("values")).split(","));
    	else
    	{
    		String[] vals = new String[columns.length];
    		
    		for(int i=0; i<columns.length; i++)
    		{
    			vals[i] = (String)data.get("column_"+columns[i].trim());
    		}
    		data.put("values", vals);
    	}
    	
    	int result = demoservice.addCommon(data);
    	
    	if(result > 0) resp.setSuccess(true);
    	else {
    		resp.setSuccess(false);
    		resp.setMessage("add fail...");
    	}
    	
        return resp;
    }

    @RequestMapping("/{table}/update")
    @ResponseBody
    public Response update(@PathVariable("table") String table, @RequestParam Map<String,Object> data) {
    	
    	Response resp = new Response();
    	
    	data.put("table", table);

//    	columns, values  처리...
//    	conds, cond_vals  처리...

    	String[] columns = ((String)data.get("columns")).split(",");
    	
    	data.put("columns", columns);
    	
    	if(data.get("values")!=null)
    		data.put("values", ((String)data.get("values")).split(","));
    	else
    	{
        	String[] values = new String[columns.length]; 
    		for(int i = 0; i<columns.length; i++) {
    			values[i] = (String)data.get("column_"+columns[i]);
    		}
    		data.put("values", values);
    	}

    	if(data.get("conds")!=null)
    	{
	    	String[] conds  = ((String)data.get("conds")).split(",");
	    	data.put("conds", conds);
	    	
	    	if(data.get("cond_vals")!=null)
	    		data.put("cond_vals", ((String)data.get("cond_vals")).split(","));
	    	else
	    	{
	    		String[] vals = new String[conds.length];
	    		
	    		for(int i=0; i<conds.length; i++)
	    		{
	    			vals[i] = (String)data.get("cond_"+conds[i].trim());
	    		}
	    		data.put("cond_vals", vals);
	    	}
	    	
    	}
    	
    	int result = demoservice.updateCommon(data);
    	
    	if(result > 0) resp.setSuccess(true);
    	else {
    		resp.setSuccess(false);
    		resp.setMessage("update fail...");
    	}
    	
        return resp;
    }

    @RequestMapping("/{table}/delete")
    @ResponseBody
    public Response delete(@PathVariable("table") String table, @RequestParam Map<String,Object> data) {
    	
    	Response resp = new Response();
    	
    	data.put("table", table);
    	
    	String conds[]  = ((String)data.get("conds")).split(",");
    	data.put("conds", conds);
    	
    	if(data.get("cond_vals")!=null)
    		data.put("cond_vals", ((String)data.get("cond_vals")).split(","));
    	else
    	{
    		String[] vals = new String[conds.length];
    		
    		for(int i=0; i<conds.length; i++)
    		{
    			vals[i] = (String)data.get("cond_"+conds[i].trim());
    		}
    		data.put("cond_vals", vals);
    	}
    	int result = demoservice.deleteCommon(data);
    	if(result > 0) resp.setSuccess(true);
    	else {
    		resp.setSuccess(false);
    		resp.setMessage("delete fail...");
    	}
        return resp;
    }	

*/
    @RequestMapping("/dong")
    @ResponseBody
    public Response donglist(@RequestParam(value = "sido", required = true) String sido) {
    	
    	Response resp = new Response();
    	    	
    	Map<String,Object> cond = new HashMap<String,Object>();
    	cond.put("sido", sido);
    	List<Map<String, Object>> lst = demoservice.listDong(cond);

    	resp.setTotal(lst.size());
    	
    	resp.setSuccess(true);
    	resp.setList((List) lst);
    	
        return resp;
    }	

    @RequestMapping("/sangho")
    @ResponseBody
    public Response sangholist(@RequestParam(value = "dong", required = true) String dong, @RequestParam(value = "upjong", required = true) String upjong) {
    	
    	Response resp = new Response();
    	    	
    	Map<String,Object> cond = new HashMap<String,Object>();
    	cond.put("dong", dong);
    	
    	String category = "";
    	String category_name = "";
    	if(upjong.equals("분식"))
    	{
        	category = "분식";
        	category_name = "상권업종중분류명";
    	}
    	else if(upjong.equals("치킨"))
    	{
        	category = "치킨 전문점";
        	category_name = "표준산업분류명";
    	}
    	else if(upjong.equals("카페"))
    	{
        	category = "커피점/카페";
        	category_name = "상권업종중분류명";
    	}
    	else if(upjong.equals("편의점"))
    	{
        	category = "체인화 편의점";
        	category_name = "표준산업분류명";
    	}
    	
    	cond.put("category", category);
    	cond.put("category_name", category_name);
    	
    	List<Map<String, Object>> lst = demoservice.listSangho(cond);

    	resp.setTotal(lst.size());
    	
    	resp.setSuccess(true);
    	resp.setList((List) lst);
    	
        return resp;
    }	
    
}

