// Generated by github.com/davyxu/tabtoy
// Version: 2.8.1
// DO NOT EDIT!!
using System.Collections.Generic;

namespace table
{
	
	// Defined in table: Globals
	public enum ActorType
	{
		
		// 唐僧
		Leader = 0, 
		
		// 孙悟空
		Monkey = 1, 
		
		// 猪八戒
		Pig = 2, 
		
		// 沙僧
		Hammer = 3, 
	
	}
	
	
	// Defined in table: Config
	public partial class Config
	{
	
		public tabtoy.Logger TableLogger = new tabtoy.Logger();
	
		
		
		public List<SampleDefine> Sample = new List<SampleDefine>(); // Sample 	
		
		public List<VerticalDefine> Vertical = new List<VerticalDefine>(); // Vertical 	
		
		public List<ExpDefine> Exp = new List<ExpDefine>(); // Exp 
	
	
		#region Index code
	 	Dictionary<long, SampleDefine> _SampleByID = new Dictionary<long, SampleDefine>();
        public SampleDefine GetSampleByID(long ID, SampleDefine def = default(SampleDefine))
        {
            SampleDefine ret;
            if ( _SampleByID.TryGetValue( ID, out ret ) )
            {
                return ret;
            }
			
			if ( def == default(SampleDefine) )
			{
				TableLogger.ErrorLine("GetSampleByID failed, ID: {0}", ID);
			}

            return def;
        }
		Dictionary<string, SampleDefine> _SampleByName = new Dictionary<string, SampleDefine>();
        public SampleDefine GetSampleByName(string Name, SampleDefine def = default(SampleDefine))
        {
            SampleDefine ret;
            if ( _SampleByName.TryGetValue( Name, out ret ) )
            {
                return ret;
            }
			
			if ( def == default(SampleDefine) )
			{
				TableLogger.ErrorLine("GetSampleByName failed, Name: {0}", Name);
			}

            return def;
        }
		Dictionary<int, ExpDefine> _ExpByLevel = new Dictionary<int, ExpDefine>();
        public ExpDefine GetExpByLevel(int Level, ExpDefine def = default(ExpDefine))
        {
            ExpDefine ret;
            if ( _ExpByLevel.TryGetValue( Level, out ret ) )
            {
                return ret;
            }
			
			if ( def == default(ExpDefine) )
			{
				TableLogger.ErrorLine("GetExpByLevel failed, Level: {0}", Level);
			}

            return def;
        }
		
	
		public VerticalDefine GetVertical( )
		{
			return Vertical[0];
		}	
	
		#endregion
		#region Deserialize code
		
		static tabtoy.DeserializeHandler<Config> ConfigDeserializeHandler = new tabtoy.DeserializeHandler<Config>(Deserialize);
		public static void Deserialize( Config ins, tabtoy.DataReader reader )
		{
			
			// Sample
			if ( reader.MatchTag(0xa0000) )
			{
				reader.ReadList_Struct<SampleDefine>( ins.Sample , SampleDefineDeserializeHandler);
			}
			
			// Vertical
			if ( reader.MatchTag(0xa0001) )
			{
				reader.ReadList_Struct<VerticalDefine>( ins.Vertical , VerticalDefineDeserializeHandler);
			}
			
			// Exp
			if ( reader.MatchTag(0xa0002) )
			{
				reader.ReadList_Struct<ExpDefine>( ins.Exp , ExpDefineDeserializeHandler);
			}
			
			
			// Build Sample Index
			for( int i = 0;i< ins.Sample.Count;i++)
			{
				var element = ins.Sample[i];
				
				ins._SampleByID.Add(element.ID, element);
				
				ins._SampleByName.Add(element.Name, element);
				
			}
			
			// Build Exp Index
			for( int i = 0;i< ins.Exp.Count;i++)
			{
				var element = ins.Exp[i];
				
				ins._ExpByLevel.Add(element.Level, element);
				
			}
			
		}
		static tabtoy.DeserializeHandler<Vec2> Vec2DeserializeHandler = new tabtoy.DeserializeHandler<Vec2>(Deserialize);
		public static void Deserialize( Vec2 ins, tabtoy.DataReader reader )
		{
			
			
			if ( reader.MatchTag(0x10000) )
			{
				ins.X = reader.ReadInt32();
			}
			
			
			if ( reader.MatchTag(0x10001) )
			{
				ins.Y = reader.ReadInt32();
			}
			
			
		}
		static tabtoy.DeserializeHandler<Prop> PropDeserializeHandler = new tabtoy.DeserializeHandler<Prop>(Deserialize);
		public static void Deserialize( Prop ins, tabtoy.DataReader reader )
		{
			
			
			if ( reader.MatchTag(0x10000) )
			{
				ins.HP = reader.ReadInt32();
			}
			
			
			if ( reader.MatchTag(0x50001) )
			{
				ins.AttackRate = reader.ReadFloat();
			}
			
			
			if ( reader.MatchTag(0x80002) )
			{
				ins.ExType = reader.ReadEnum<ActorType>();
			}
			
			
		}
		static tabtoy.DeserializeHandler<SampleDefine> SampleDefineDeserializeHandler = new tabtoy.DeserializeHandler<SampleDefine>(Deserialize);
		public static void Deserialize( SampleDefine ins, tabtoy.DataReader reader )
		{
			
			// 唯一ID
			if ( reader.MatchTag(0x20000) )
			{
				ins.ID = reader.ReadInt64();
			}
			
			// 名称
			if ( reader.MatchTag(0x60001) )
			{
				ins.Name = reader.ReadString();
			}
			
			// 图标ID
			if ( reader.MatchTag(0x10002) )
			{
				ins.IconID = reader.ReadInt32();
			}
			
			// 攻击率
			if ( reader.MatchTag(0x50003) )
			{
				ins.NumericalRate = reader.ReadFloat();
			}
			
			// 物品id
			if ( reader.MatchTag(0x10004) )
			{
				ins.ItemID = reader.ReadInt32();
			}
			
			// BuffID
			if ( reader.MatchTag(0x10005) )
			{
				reader.ReadList_Int32( ins.BuffID );
			}
			
			
			if ( reader.MatchTag(0x90006) )
			{
				ins.Pos = reader.ReadStruct<Vec2>(Vec2DeserializeHandler);
			}
			
			// 类型
			if ( reader.MatchTag(0x80007) )
			{
				ins.Type = reader.ReadEnum<ActorType>();
			}
			
			// 技能ID列表
			if ( reader.MatchTag(0x10008) )
			{
				reader.ReadList_Int32( ins.SkillID );
			}
			
			// 单结构解析
			if ( reader.MatchTag(0x90009) )
			{
				ins.SingleStruct = reader.ReadStruct<Prop>(PropDeserializeHandler);
			}
			
			// 字符串结构
			if ( reader.MatchTag(0x9000a) )
			{
				reader.ReadList_Struct<Prop>( ins.StrStruct , PropDeserializeHandler);
			}
			
			
		}
		static tabtoy.DeserializeHandler<PeerData> PeerDataDeserializeHandler = new tabtoy.DeserializeHandler<PeerData>(Deserialize);
		public static void Deserialize( PeerData ins, tabtoy.DataReader reader )
		{
			
			
			if ( reader.MatchTag(0x60000) )
			{
				ins.Name = reader.ReadString();
			}
			
			
			if ( reader.MatchTag(0x60001) )
			{
				ins.Type = reader.ReadString();
			}
			
			
		}
		static tabtoy.DeserializeHandler<VerticalDefine> VerticalDefineDeserializeHandler = new tabtoy.DeserializeHandler<VerticalDefine>(Deserialize);
		public static void Deserialize( VerticalDefine ins, tabtoy.DataReader reader )
		{
			
			// 服务器IP
			if ( reader.MatchTag(0x60000) )
			{
				ins.ServerIP = reader.ReadString();
			}
			
			// 调试模式
			if ( reader.MatchTag(0x70001) )
			{
				ins.DebugMode = reader.ReadBool();
			}
			
			// 客户端人数限制
			if ( reader.MatchTag(0x10002) )
			{
				ins.ClientLimit = reader.ReadInt32();
			}
			
			// 端
			if ( reader.MatchTag(0x90003) )
			{
				ins.Peer = reader.ReadStruct<PeerData>(PeerDataDeserializeHandler);
			}
			
			
			if ( reader.MatchTag(0x50004) )
			{
				ins.Float = reader.ReadFloat();
			}
			
			
		}
		static tabtoy.DeserializeHandler<ExpDefine> ExpDefineDeserializeHandler = new tabtoy.DeserializeHandler<ExpDefine>(Deserialize);
		public static void Deserialize( ExpDefine ins, tabtoy.DataReader reader )
		{
			
			// 唯一ID
			if ( reader.MatchTag(0x10000) )
			{
				ins.Level = reader.ReadInt32();
			}
			
			// 经验值
			if ( reader.MatchTag(0x10001) )
			{
				ins.Exp = reader.ReadInt32();
			}
			
			// 布尔检查
			if ( reader.MatchTag(0x70002) )
			{
				ins.BoolChecker = reader.ReadBool();
			}
			
			// 类型
			if ( reader.MatchTag(0x80003) )
			{
				ins.Type = reader.ReadEnum<ActorType>();
			}
			
			
		}
		#endregion
	

	} 
	// Defined in table: Globals
	public partial class Vec2
	{
	
		
		
		public int X = 0;  	
		
		public int Y = 0;  
	
	

	} 
	// Defined in table: Sample
	public partial class Prop
	{
	
		
		// 血量
		public int HP = 10;  	
		// 攻击速率
		public float AttackRate = 0f;  	
		// 额外类型
		public ActorType ExType = ActorType.Leader;  
	
	

	} 
	// Defined in table: Sample
	public partial class SampleDefine
	{
	
		
		
		public long ID = 0; // 唯一ID 	
		
		public string Name = ""; // 名称 	
		
		public int IconID = 0; // 图标ID 	
		
		public float NumericalRate = 0f; // 攻击率 	
		
		public int ItemID = 100; // 物品id 	
		
		public List<int> BuffID = new List<int>(); // BuffID 	
		
		public Vec2 Pos = new Vec2();  	
		
		public ActorType Type = ActorType.Leader; // 类型 	
		
		public List<int> SkillID = new List<int>(); // 技能ID列表 	
		
		public Prop SingleStruct = new Prop(); // 单结构解析 	
		
		public List<Prop> StrStruct = new List<Prop>(); // 字符串结构 
	
	

	} 
	// Defined in table: Vertical
	public partial class PeerData
	{
	
		
		// 名字
		public string Name = "";  	
		// 类型
		public string Type = "";  
	
	

	} 
	// Defined in table: Vertical
	public partial class VerticalDefine
	{
	
		
		
		public string ServerIP = ""; // 服务器IP 	
		
		public bool DebugMode = false; // 调试模式 	
		
		public int ClientLimit = 0; // 客户端人数限制 	
		
		public PeerData Peer = new PeerData(); // 端 	
		
		public float Float = 0.5f;  
	
	

	} 
	// Defined in table: Exp
	public partial class ExpDefine
	{
	
		
		
		public int Level = 0; // 唯一ID 	
		
		public int Exp = 0; // 经验值 	
		
		public bool BoolChecker = false; // 布尔检查 	
		
		public ActorType Type = ActorType.Leader; // 类型 
	
	

	} 

}
