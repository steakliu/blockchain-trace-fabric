package com.ruoyi.traces.service;

import com.ruoyi.traces.domain.TraceTask;
import com.ruoyi.traces.domain.vo.CropsVo;

import java.util.List;

public interface ProductService {
    List<CropsVo> queryCropsList(Integer deptId);

    int addTask(TraceTask traceTask);

    int updateMachingStatus(CropsVo cropsVo);

    TraceTask queryTaskByCropsId(String cropsId);

    int updateProductWriteStatus(String cropsId);

    List<TraceTask> queryTaskList();

    int productOutFactory(String cropsId);
}
