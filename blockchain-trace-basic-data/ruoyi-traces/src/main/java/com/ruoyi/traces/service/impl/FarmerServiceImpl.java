package com.ruoyi.traces.service.impl;

import com.ruoyi.common.core.domain.entity.SysDept;
import com.ruoyi.common.core.domain.entity.SysUser;
import com.ruoyi.system.mapper.SysDeptMapper;
import com.ruoyi.system.mapper.SysUserMapper;
import com.ruoyi.traces.domain.TraceTransport;
import com.ruoyi.traces.domain.TraceUserCrops;
import com.ruoyi.traces.mapper.TraceTransportMapper;
import com.ruoyi.traces.mapper.TraceUserCropsMapper;
import com.ruoyi.traces.service.FarmerService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class FarmerServiceImpl implements FarmerService {
    @Autowired
    private TraceUserCropsMapper traceCropsMapper;
    @Autowired
    private SysUserMapper sysUserMapper;
    @Autowired
    private TraceTransportMapper traceTransportMapper;
    @Autowired
    private SysDeptMapper sysDeptMapper;

    @Override
    public int saveCrops(TraceUserCrops traceCrops) {
        return traceCropsMapper.insert(traceCrops);
    }

    @Override
    public List<TraceUserCrops> getCropsByUsername(String username) {
        return traceCropsMapper.getCropsByUsername(username);
    }

    @Override
    public List<SysUser> getAllDriverByDeptId(String deptId) {
        return sysUserMapper.selectUserByDeptId(deptId);
    }

    @Override
    public int addTransport(TraceTransport traceTransport) {
        int i = traceCropsMapper.updateByCropsId(traceTransport.getCropsId());
        return traceTransportMapper.insertSelective(traceTransport);
    }

    @Override
    public List<SysDept> getFactoryByDeptId(Integer deptId) {
        return sysDeptMapper.getFactoryByDeptId(deptId);
    }


}
