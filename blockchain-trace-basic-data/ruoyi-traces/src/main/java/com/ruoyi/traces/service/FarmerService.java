package com.ruoyi.traces.service;

import com.ruoyi.common.core.domain.entity.SysDept;
import com.ruoyi.common.core.domain.entity.SysUser;
import com.ruoyi.traces.domain.TraceTransport;
import com.ruoyi.traces.domain.TraceUserCrops;

import java.util.List;

public interface FarmerService {
    int saveCrops(TraceUserCrops traceCrops);

    List<TraceUserCrops> getCropsByUsername(String username);

    List<SysUser> getAllDriverByDeptId(String deptId);

    int addTransport(TraceTransport traceTransport);

    List<SysDept> getFactoryByDeptId(Integer deptId);
}
