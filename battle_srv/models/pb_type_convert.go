package models

import (
	pb "battle_srv/protos"
	"jsexport/battle"
)

func toPbRoomDownsyncFrame(rdf *battle.RoomDownsyncFrame) *pb.RoomDownsyncFrame {
	if nil == rdf {
		return nil
	}
	ret := &pb.RoomDownsyncFrame{
		Id:                     rdf.Id,
		PlayersArr:             make([]*pb.PlayerDownsync, len(rdf.PlayersArr), len(rdf.PlayersArr)),
		BulletLocalIdCounter:   rdf.BulletLocalIdCounter,
		MeleeBullets:           make([]*pb.MeleeBullet, len(rdf.MeleeBullets), len(rdf.MeleeBullets)),
		FireballBullets:        make([]*pb.FireballBullet, len(rdf.FireballBullets), len(rdf.FireballBullets)),
		CountdownNanos:         rdf.CountdownNanos,
		BackendUnconfirmedMask: rdf.BackendUnconfirmedMask,
		ShouldForceResync:      rdf.ShouldForceResync,
	}

	for i, last := range rdf.PlayersArr {
		pbPlayer := &pb.PlayerDownsync{
			Id:                last.Id,
			VirtualGridX:      last.VirtualGridX,
			VirtualGridY:      last.VirtualGridY,
			DirX:              last.DirX,
			DirY:              last.DirY,
			VelX:              last.VelX,
			VelY:              last.VelY,
			FramesToRecover:   last.FramesToRecover,
			FramesInChState:   last.FramesInChState,
			ActiveSkillId:     last.ActiveSkillId,
			ActiveSkillHit:    last.ActiveSkillHit,
			FramesInvinsible:  last.FramesInvinsible,
			Speed:             last.Speed,
			BattleState:       last.BattleState,
			CharacterState:    last.CharacterState,
			InAir:             last.InAir,
			JoinIndex:         last.JoinIndex,
			BulletTeamId:      last.BulletTeamId,
			ChCollisionTeamId: last.ChCollisionTeamId,
			Hp:                last.Hp,
			MaxHp:             last.MaxHp,
			ColliderRadius:    last.ColliderRadius,
			Score:             last.Score,
			Removed:           last.Removed,
		}
		ret.PlayersArr[i] = pbPlayer
	}

	for i, last := range rdf.MeleeBullets {
		pbBullet := &pb.MeleeBullet{
			BulletLocalId:           last.BulletLocalId,
			OriginatedRenderFrameId: last.OriginatedRenderFrameId,
			OffenderJoinIndex:       last.OffenderJoinIndex,

			StartupFrames:      last.StartupFrames,
			CancellableStFrame: last.CancellableStFrame,
			CancellableEdFrame: last.CancellableEdFrame,
			ActiveFrames:       last.ActiveFrames,

			HitStunFrames:   last.HitStunFrames,
			BlockStunFrames: last.BlockStunFrames,
			PushbackVelX:    last.PushbackVelX,
			PushbackVelY:    last.PushbackVelY,
			Damage:          last.Damage,

			SelfLockVelX: last.SelfLockVelX,
			SelfLockVelY: last.SelfLockVelY,

			HitboxOffsetX: last.HitboxOffsetX,
			HitboxOffsetY: last.HitboxOffsetY,
			HitboxSizeX:   last.HitboxSizeX,
			HitboxSizeY:   last.HitboxSizeY,

			BlowUp: last.BlowUp,
			TeamId: last.TeamId,
		}
		ret.MeleeBullets[i] = pbBullet
	}

	for i, last := range rdf.FireballBullets {
		pbBullet := &pb.FireballBullet{
			BulletLocalId:           last.BulletLocalId,
			OriginatedRenderFrameId: last.OriginatedRenderFrameId,
			OffenderJoinIndex:       last.OffenderJoinIndex,

			StartupFrames:      last.StartupFrames,
			CancellableStFrame: last.CancellableStFrame,
			CancellableEdFrame: last.CancellableEdFrame,
			ActiveFrames:       last.ActiveFrames,

			HitStunFrames:   last.HitStunFrames,
			BlockStunFrames: last.BlockStunFrames,
			PushbackVelX:    last.PushbackVelX,
			PushbackVelY:    last.PushbackVelY,
			Damage:          last.Damage,

			SelfLockVelX: last.SelfLockVelX,
			SelfLockVelY: last.SelfLockVelY,

			HitboxOffsetX: last.HitboxOffsetX,
			HitboxOffsetY: last.HitboxOffsetY,
			HitboxSizeX:   last.HitboxSizeX,
			HitboxSizeY:   last.HitboxSizeY,

			BlowUp: last.BlowUp,
			TeamId: last.TeamId,

			VirtualGridX: last.VirtualGridX,
			VirtualGridY: last.VirtualGridY,
			DirX:         last.DirX,
			DirY:         last.DirY,
			VelX:         last.VelX,
			VelY:         last.VelY,
			Speed:        last.Speed,
		}
		ret.FireballBullets[i] = pbBullet
	}

	return ret
}

func toPbPlayers(modelInstances map[int32]*Player, withMetaInfo bool) []*pb.PlayerDownsync {
	toRet := make([]*pb.PlayerDownsync, len(modelInstances), len(modelInstances))
	if nil == modelInstances {
		return toRet
	}

	for _, last := range modelInstances {
		pbPlayer := &pb.PlayerDownsync{
			Id:                last.Id,
			VirtualGridX:      last.VirtualGridX,
			VirtualGridY:      last.VirtualGridY,
			DirX:              last.DirX,
			DirY:              last.DirY,
			VelX:              last.VelX,
			VelY:              last.VelY,
			FramesToRecover:   last.FramesToRecover,
			FramesInChState:   last.FramesInChState,
			ActiveSkillId:     last.ActiveSkillId,
			ActiveSkillHit:    last.ActiveSkillHit,
			FramesInvinsible:  last.FramesInvinsible,
			Speed:             last.Speed,
			BattleState:       last.BattleState,
			CharacterState:    last.CharacterState,
			InAir:             last.InAir,
			JoinIndex:         last.JoinIndex,
			BulletTeamId:      last.BulletTeamId,
			ChCollisionTeamId: last.ChCollisionTeamId,
			ColliderRadius:    last.ColliderRadius,
			Score:             last.Score,
			Removed:           last.Removed,
		}
		if withMetaInfo {
			pbPlayer.Name = last.Name
			pbPlayer.DisplayName = last.DisplayName
			pbPlayer.Avatar = last.Avatar
		}
		toRet[last.JoinIndex-1] = pbPlayer
	}

	return toRet
}

func toJsPlayers(modelInstances map[int32]*Player) []*battle.PlayerDownsync {
	toRet := make([]*battle.PlayerDownsync, len(modelInstances), len(modelInstances))
	if nil == modelInstances {
		return toRet
	}

	for _, last := range modelInstances {
		toRet[last.JoinIndex-1] = &battle.PlayerDownsync{
			Id:                last.Id,
			VirtualGridX:      last.VirtualGridX,
			VirtualGridY:      last.VirtualGridY,
			DirX:              last.DirX,
			DirY:              last.DirY,
			VelX:              last.VelX,
			VelY:              last.VelY,
			FramesToRecover:   last.FramesToRecover,
			FramesInChState:   last.FramesInChState,
			ActiveSkillId:     last.ActiveSkillId,
			ActiveSkillHit:    last.ActiveSkillHit,
			FramesInvinsible:  last.FramesInvinsible,
			Speed:             last.Speed,
			BattleState:       last.BattleState,
			CharacterState:    last.CharacterState,
			JoinIndex:         last.JoinIndex,
			BulletTeamId:      last.BulletTeamId,
			ChCollisionTeamId: last.ChCollisionTeamId,
			Hp:                last.Hp,
			MaxHp:             last.MaxHp,
			ColliderRadius:    last.ColliderRadius,
			InAir:             last.InAir,
			Score:             last.Score,
			Removed:           last.Removed,
		}
	}

	return toRet
}
